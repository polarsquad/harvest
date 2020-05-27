package harvest

import (
	"log"
	"testing"

	"encoding/json"
	"io/ioutil"
	"time"
)

var (
	testDataFile = "test/test_timeEntries.json"
	entries      = []float64{7.5, 5.5, 9.5, 6.5, 1, 7.5, 5.5, 4, 7.5, 2.5, 2, 3, 0.5, 1, 1, 2, 1.5, 1.5, 7.5, 2, 7.5}
)

func TestDailyHours(t *testing.T) {
	var expect float64
	expect = 7.5

	// var c config.Config
	// h := Init(&c)
	body, err := ioutil.ReadFile(testDataFile)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}

	var timeEntries TimeEntries
	json.Unmarshal(body, &timeEntries)

	date, _ := time.Parse("2006-01-02", "2019-04-15")
	result := timeEntries.DailyHours(date)
	if result != expect {
		t.Errorf("DailyHours: expecting %v, got %v!", expect, result)
	} else {
		t.Logf("[PASS] DailyHours: expecting %v, got %v!", expect, result)
	}

	expect = 5.5
	date, _ = time.Parse("2006-01-02", "2019-04-16")
	result = timeEntries.DailyHours(date)
	if result != expect {
		t.Errorf("DailyHours: expecting %v, got %v!", expect, result)
	} else {
		t.Logf("[PASS] DailyHours: expecting %v, got %v!", expect, result)
	}
}

func TestDailyTotals(t *testing.T) {
	var expectHours, expectSaldo, expectOvertime float64
	expectHours = 5.5
	expectSaldo = 0.0
	expectOvertime = -2.0

	// var c config.Config
	// h := Init(&c)
	body, _ := ioutil.ReadFile(testDataFile)
	// if err != nil {
	// 	log.Fatalf(err.Error)
	// }

	var timeEntries TimeEntries
	json.Unmarshal(body, &timeEntries)

	date, _ := time.Parse("2006-01-02", "2019-04-16")
	resultHours, resultSaldo, resultOvertime := timeEntries.DailyTotals(date)
	if resultHours != expectHours {
		t.Errorf("DailyTotals.Hours: expecting %v, got %v!", expectHours, resultHours)
	} else {
		t.Logf("[PASS] DailyTotals.Hours: expecting %v, got %v!", expectHours, resultHours)
	}

	if resultSaldo != expectSaldo {
		t.Errorf("DailyTotals.Saldo: expecting %v, got %v!", expectSaldo, resultSaldo)
	} else {
		t.Logf("[PASS] DailyTotals.Saldo: expecting %v, got %v!", expectSaldo, resultSaldo)
	}
	if resultOvertime != expectOvertime {
		t.Errorf("DailyTotals.Overtime: expecting %v, got %v!", expectOvertime, resultOvertime)
	} else {
		t.Logf("[PASS] DailyTotals.Overtime: expecting %v, got %v!", expectOvertime, resultOvertime)
	}
}

func TestTotalHours(t *testing.T) {
	var expect float64
	// weekOne := []string{"2019-04-15", "2019-04-21"}
	// weekTwo := []string{"2019-04-22", "2019-04-28"}
	expectWeekOne := 47.0
	expectWeekTwo := 39.5
	expect = expectWeekOne + expectWeekTwo

	// var c config.Config
	// h := Init(&c)

	body, err := ioutil.ReadFile(testDataFile)
	if err != nil {
		log.Fatalf("[ERROR] Could not read test data.\n[ERROR] %v", err)
	}

	var timeEntries TimeEntries
	json.Unmarshal(body, &timeEntries)

	// date, _ := time.Parse("2006-01-02", "2019-02-28")
	result := timeEntries.TotalHours()
	if result != expect {
		t.Errorf("TotalHours: expecting %v, got %v!", expect, result)
	} else {
		t.Logf("[PASS] TotalHours: expecting %v, got %v!", expect, result)
	}
}
