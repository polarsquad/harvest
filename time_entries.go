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

func addParamsToURL(baseURL string, opt interface{}) (string, error) { //TODO: Move to helpers.go
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

func (h *Harvest) getURL(method string, url string) ([]byte, error) { //TODO: Move to heplers.go
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

// GetEntries fetches all the TimeEntries with the provided timespan.
// If from and/or to dates are not defined, it will fetch all the TimeEntries.
//
// Parameters for the function:
// from: time.Time with format "2006-01-02"
// to: time.Time with format "2006-01-02"
// u: User, specifies which users TimeEntries are fetched.
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
	log.Printf("Fetching URL: %s", l.Next)
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
