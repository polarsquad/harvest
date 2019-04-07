package harvest

import (
	// "fmt"
	"log"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

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

// func (h *Harvest) getURL(method string, url string) ([]byte, error) {
// 	Client := &http.Client{}

// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Set("User-Agent", "Harvest Slack Bot")
// 	req.Header.Set("Harvest-Account-ID", h.API.AccountID)
// 	req.Header.Set("Authorization", "Bearer "+h.API.AuthToken)
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := Client.Do(req)
// 	body, err := ioutil.ReadAll(resp.Body)
// 	defer resp.Body.Close()

// 	// var jsonResponse map[string]interface{}

// 	// json.Unmarshal(body, &jsonResponse)

// 	return body, err
// }

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

// func getAllEntries(l *structs.Links) *TimeEntries {

// 	return nil
// }

func (h *Harvest) GetEntries(from time.Time, to time.Time, u User) *TimeEntries {
	// Start with fetching the entries
	url := "https://api.harvestapp.com/v2/time_entries"

	// Let's build the URL with parameters.
	params := GetTimeEntriesParams{
		UserID:  int64(u.ID),
		From:    from.Format("2006-01-02"),
		PerPage: 10,
		// IsBilled: true,
		To: to.Format("2006-01-02"),
	}

	urlWithParams, _ := addParamsToURL(url, &params)
	body, err := h.getURL("GET", urlWithParams)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var times TimeEntries
	json.Unmarshal(body, &times)
	log.Printf("TotalEntries: %v, Pages: %v\n", times.TotalEntries, times.TotalPages)

	// IF: returned entries have Next link, then we fetch the additional pages
	if times.Links.Next != "" {
		// Then call getAllEntries
		// time.Sleep(1 * time.Second) // This might be needed, if Harvest API starts throttling
		entries := h.getAllEntries(times.Links)
		times.Entries = append(times.Entries, entries...)

	}

	return &times
}

// func (h *Harvest) getAllEntries(l structs.Links, entries *[]structs.Entries, i int) Entries {
func (h *Harvest) getAllEntries(l structs.Links) Entries {
	// getURL to fetch additional Entries from Links.Next URL
	body, _ := h.getURL("GET", l.Next)
	var times TimeEntries
	json.Unmarshal(body, &times)
	if times.Links.Next != "" {
		// IF Next URL still available, let's call getAllEntries() again.
		entries := h.getAllEntries(times.Links)
		times.Entries = append(times.Entries, entries...)
	} else {
		// Return times.Entries since no Links.Next URL available
		return times.Entries
	}

	return times.Entries
}

// func (e *structs.TimeEntries) totalHours() float64 {
// 	var hours float64
// 	for _, v := range e.Entries {
// 		hours = hours + v.Hours
// 	}
// 	// fmt.Printf("Total hours: %v\n", hours)
// 	return hours
// }

func (h *Harvest) TotalHours() float64 {
	var hours float64
	for _, v := range h.TimeEntries.Entries {
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

// func (h *Harvest) GetHours(u *User) float64 {
// 	// url := "https://api.harvestapp.com/v2/time_entries?from=2019-02-20"
// 	url := "https://api.harvestapp.com/v2/time_entries"
// 	// Client := &http.Client{}

// 	// Let's build the URL with parameters.
// 	params := GetTimeEntriesParams{
// 		UserID:  int64(u.ID),
// 		From:    "2019-01-01",
// 		PerPage: 20,
// 		// IsBilled: true,
// 		To: "2019-01-31",
// 	}
// 	// prms := map[string]interface{}
// 	// prms = params
// 	// fmt.Printf("OPTS: %v", params)

// 	urlWithParams, _ := addParamsToURL(url, &params)
// 	log.Printf("URL: %v\n", urlWithParams)

// 	// req, _ := http.NewRequest("GET", urlWithParams, nil)
// 	// req.Header.Set("User-Agent", "Go Harvest API Sample")
// 	// req.Header.Set("Harvest-Account-ID", AccountId)
// 	// req.Header.Set("Authorization", "Bearer "+AuthToken)
// 	// req.Header.Set("Content-Type", "application/json")

// 	// resp, _ := Client.Do(req)
// 	body, _ := h.getURL("GET", urlWithParams)
// 	// body, _ := ioutil.ReadAll(resp.Body)
// 	// defer resp.Body.Close()

// 	var times TimeEntries

// 	json.Unmarshal(body, &times)
// 	log.Printf("TOTAL ENTRIES: %v\n", times.TotalEntries)
// 	log.Printf("TOTAL PAGES: %v\n", times.TotalPages)
// 	log.Printf("START ENTRIES: %v\n", len(times.Entries))

// 	// var allEntries []structs.Entries
// 	var allEntries []structs.Entries

// 	i := 1
// 	// getAll := times.Links.getAllEntries(&allEntries, i)
// 	getAll := h.getAllEntries(times.Links, &allEntries, i) // TODO: CHANGE REFERENCING
// 	log.Printf("getAll: %v\n", len(getAll))
// 	for _, v := range getAll {
// 		allEntries = append(allEntries, v)
// 	}
// 	for _, v := range times.Entries {
// 		allEntries = append(allEntries, v)
// 	}
// 	log.Printf("ALLENTRIES LEN: %v\n", len(allEntries))

// 	// for _, v := range times.Entries {
// 	// 	// fmt.Printf("ORIG DATE: %v\n", v.SpentDate)
// 	// 	date, _ := time.Parse("2006-01-02", v.SpentDate)
// 	// 	// fmt.Printf("DATE: %v\n", date)
// 	// 	fmt.Printf("%v %v, Hours: %v\n", date.Weekday(), date.Format("01.02.2006"), v.Hours)
// 	// }

// 	times.Entries = []structs.Entries(allEntries)
// 	// log.Printf("Total Hours: %v\n", times.totalHours())
// 	day := "2019-02-20"
// 	log.Printf("Hours on date %s: %v", day, times.dailyHours(day))

// 	return times.totalHours()
// }
