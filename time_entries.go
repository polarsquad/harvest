package harvest

import (
	// "fmt"
	"log"

	"encoding/json"
	"time"

	// "github.com/google/go-querystring/query"
	"github.com/polarsquad/harvest/structs"
)

// GetEntries fetches all the TimeEntries with the provided timespan.
// If from and/or to dates are not defined, it will fetch all the TimeEntries.
//
// Parameters for the function:
// from: time.Time with format "2006-01-02"
// to: time.Time with format "2006-01-02"
// u: User, specifies which users TimeEntries are fetched.
func (h *Harvest) GetEntries(from time.Time, to time.Time, u *User) *TimeEntries {
	// Let's build the URL with parameters.
	params := GetTimeEntriesParams{
		UserID:  int64(u.ID),
		From:    from.Format("2006-01-02"),
		PerPage: 10,
		// IsBilled: true,
		To: to.Format("2006-01-02"),
	}

	urlWithParams, _ := addParamsToURL(timeEntriesURL, &params)
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
