package harvest

import (
	"fmt"
	// "log"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/polarsquad/harvest/structs"
)

func addParamsToURL(baseURL string, opt interface{}) (string, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return baseURL, err
	}

	vs, err := query.Values(opt)
	if err != nil {
		return baseURL, err
	}

	url.RawQuery = vs.Encode()
	return url.String(), nil
}

func (h *Harvest) getURL(method string, url string) ([]byte, error) {
	Client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Harvest Slack Bot")
	req.Header.Set("Harvest-Account-ID", h.API.AccountID)
	req.Header.Set("Authorization", "Bearer "+h.API.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := Client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// var jsonResponse map[string]interface{}

	// json.Unmarshal(body, &jsonResponse)

	return body, err
}

func (h *Harvest) getAllEntries(l structs.Links, entries *[]structs.Entries, i int) []structs.Entries {
	fmt.Println("---- START ----")
	fmt.Printf("entries: %v\n", len(*entries))
	fmt.Printf("NEXT LINK: %v\n", l.Next)
	// i++
	// entry := *entries
	var entry []structs.Entries
	// var entr Entries
	var times structs.TimeEntries
	body, _ := h.getURL("GET", l.Next)
	json.Unmarshal(body, &times)
	if times.Links.Next != "" {
		fmt.Printf("    IF ENTRIES: %v\n", len(times.Entries))
		// body, _ := getURL("GET", times.Links.Next)
		// json.Unmarshal(body, &times)
		i++
		entry = []structs.Entries(h.getAllEntries(l, &entry, i))
		for _, v := range times.Entries {
			entry = append(entry, v)
		}
	} else {
		fmt.Printf("    ELSE ENTRIES (%v): %v\n", i, len(times.Entries))
		for _, v := range times.Entries {
			entry = append(entry, v)
		}
		for _, v := range *entries {
			entry = append(entry, v)
		}
		fmt.Printf("(%v)ENTRY LEN: %v\n", i, len(entry))
		i--
		return entry
	}
	// i--
	fmt.Printf("(%v) RIGHT AFTER IF: %v\n", i, len(entry))
	// ** THIS IS ORIGINAL FUNCTIONING...partially
	// if l.Next != "" {
	// 	body, _ := getURL("GET", l.Next)
	// 	json.Unmarshal(body, &times)
	// 	entry = []Entries(times.Links.getAllEntries(&entry))
	// } else {
	// 	body, _ := getURL("GET", l.Next)
	// 	json.Unmarshal(body, &times)
	// 	entry = []Entries(times.Links.getAllEntries(&entry))
	// }

	for _, v := range *entries {

		entry = append(entry, v)
	}
	fmt.Printf("(%v) ENTRY LEN: %v\n", i, len(entry))
	// entry = append(entry, times.Entries)
	// entry = []Entries(times.Entries)
	fmt.Printf("NEXT: %v\n", l.Next)
	// fmt.Printf("Entries: %v\n", len(times.Entries))
	fmt.Println("---- END ----")
	i--
	return entry
}

func (e *TimeEntries) totalHours() float64 {
	var hours float64
	for _, v := range e.Entries {
		hours = hours + v.Hours
	}
	// fmt.Printf("Total hours: %v\n", hours)
	return hours
}

func (e *TimeEntries) dailyHours(daySelector string) float64 {
	var selection Entries

	for i := 0; i < len(e.Entries); i++ {
		if e.Entries[i].SpentDate == daySelector {
			selection = append(selection, e.Entries[i])
		}
	}

	dayHours := selection.dayTotal()

	// for _, v := range e.Entries {
	// 	if v.SpentDate == daySelector {
	// 		dayHours = dayHours + v.Hours
	// 	}
	// }
	return dayHours
}

func (e Entries) dayTotal() float64 {
	var hours float64
	for _, v := range e {
		hours = hours + v.Hours
	}
	return hours
}

func (h *Harvest) GetHours(u structs.User) {
	// url := "https://api.harvestapp.com/v2/time_entries?from=2019-02-20"
	url := "https://api.harvestapp.com/v2/time_entries"
	// Client := &http.Client{}

	// Let's build the URL with parameters.
	params := GetTimeEntriesParams{
		UserID:  int64(u.ID),
		From:    "2019-01-01",
		PerPage: 20,
		// IsBilled: true,
		To: "2019-01-31",
	}
	// prms := map[string]interface{}
	// prms = params
	// fmt.Printf("OPTS: %v", params)

	urlWithParams, _ := addParamsToURL(url, &params)
	fmt.Printf("URL: %v\n", urlWithParams)

	// req, _ := http.NewRequest("GET", urlWithParams, nil)
	// req.Header.Set("User-Agent", "Go Harvest API Sample")
	// req.Header.Set("Harvest-Account-ID", AccountId)
	// req.Header.Set("Authorization", "Bearer "+AuthToken)
	// req.Header.Set("Content-Type", "application/json")

	// resp, _ := Client.Do(req)
	body, _ := h.getURL("GET", urlWithParams)
	// body, _ := ioutil.ReadAll(resp.Body)
	// defer resp.Body.Close()

	var times TimeEntries

	json.Unmarshal(body, &times)
	fmt.Printf("TOTAL ENTRIES: %v\n", times.TotalEntries)
	fmt.Printf("TOTAL PAGES: %v\n", times.TotalPages)
	fmt.Printf("START ENTRIES: %v\n", len(times.Entries))

	// var allEntries []structs.Entries
	var allEntries []structs.Entries

	i := 1
	// getAll := times.Links.getAllEntries(&allEntries, i)
	getAll := h.getAllEntries(times.Links, &allEntries, i) // TODO: CHANGE REFERENCING
	fmt.Printf("getAll: %v\n", len(getAll))
	for _, v := range getAll {
		allEntries = append(allEntries, v)
	}
	for _, v := range times.Entries {
		allEntries = append(allEntries, v)
	}
	fmt.Printf("ALLENTRIES LEN: %v\n", len(allEntries))

	// for _, v := range times.Entries {
	// 	// fmt.Printf("ORIG DATE: %v\n", v.SpentDate)
	// 	date, _ := time.Parse("2006-01-02", v.SpentDate)
	// 	// fmt.Printf("DATE: %v\n", date)
	// 	fmt.Printf("%v %v, Hours: %v\n", date.Weekday(), date.Format("01.02.2006"), v.Hours)
	// }

	times.Entries = []structs.Entries(allEntries)
	fmt.Printf("Total Hours: %v\n", times.totalHours())
	day := "2019-02-20"
	fmt.Printf("Hours on date %s: %v", day, times.dailyHours(day))
}
