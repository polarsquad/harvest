package harvest

import (
	"testing"

	"encoding/json"
	"github.com/polarsquad/harvest/config"
	"io/ioutil"
	"time"
)

func TestDailyHours(t *testing.T) {
	var expect float64
	expect = 7.5

	var c config.Config
	h := Init(&c)
	body, _ := ioutil.ReadFile("test/sample_timeEntries.json")
	// if err != nil {
	// 	log.Fatalf(err.Error)
	// }

	// var times structs.TimeEntries
	json.Unmarshal(body, &h.TimeEntries)

	date, _ := time.Parse("2006-01-02", "2019-02-28")
	result := h.TimeEntries.DailyHours(date)
	if result != expect {
		t.Errorf("DailyHours: expecting %v, got %v!", expect, result)
	}
	// log.Printf("DayliHours: expecting %v, got %v!", expect, result)
}

func TestTotalHours(t *testing.T) {
	var expect float64
	expect = 53.5

	var c config.Config
	h := Init(&c)
	body, _ := ioutil.ReadFile("test/sample_timeEntries.json")
	// if err != nil {
	// 	log.Fatalf(err.Error)
	// }

	// var times structs.TimeEntries
	json.Unmarshal(body, &h.TimeEntries)

	// date, _ := time.Parse("2006-01-02", "2019-02-28")
	result := h.TotalHours()
	if result != expect {
		t.Errorf("TotalHours: expecting %v, got %v!", expect, result)
	}
	// log.Printf("DayliHours: expecting %v, got %v!", expect, result)
}
