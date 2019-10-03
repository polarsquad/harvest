package harvest

import (
	"github.com/polarsquad/harvest/structs"
	_ "log"
	"time"
)

// GetOvertime counts overtime hours from TimeEntries, using also dayTotal function as a helper.
func (e *TimeEntries) GetOvertime(from time.Time, to time.Time) (totalOvertime float64) {
	for d := from; d.Before(to) || d.Equal(to); d = d.AddDate(0, 0, 1) {
		// dailyHours := e.DailyHours(d)
		_, saldo, overtime := e.DailyTotals(d)

		totalOvertime = totalOvertime + overtime
		if saldo != 0 {
			totalOvertime = totalOvertime - saldo
		}

		// if IsWorkday(d) {
		// 	if dailyHours > 7.5 {
		// 		overtime = overtime + dailyHours - 7.5
		// 	}
		// 	if dailyHours < 7.5 {
		// 		overtime = overtime + dailyHours - 7.5
		// 	}
		// } else {
		// 	overtime = overtime + dailyHours
		// }
	}
	return
}

// TotalHours counts total logged hours from the TimeEntries struct
func (h *Harvest) TotalHours() float64 {
	var hours float64
	for _, v := range h.TimeEntries.Entries {
		hours = hours + v.Hours
	}
	// fmt.Printf("Total hours: %v\n", hours)
	return hours
}

// Total counts total logged hours from the TimeEntries struct
func (e TimeEntries) Total() float64 {
	var hours float64
	for _, v := range e.Entries {
		hours = hours + v.Hours
	}
	// fmt.Printf("Total hours: %v\n", hours)
	return hours
}

// DailyTotals counts total logged in hours for selected date.
// Needs daySelector(time.Time) as parameter for selected date.
// Will output hours, saldo and overtime.
// Hours is logged hours excluding spent flexitime.
// Saldo is spent flexitime for that day.
// Overtime is overtime for that day.
func (e *TimeEntries) DailyTotals(daySelector time.Time) (hours float64, saldo float64, overtime float64) {
	// var selection Entries
	// var saldoused = false

	date := daySelector.Format("2006-01-02") // TODO: Switch to get formatter from config

	for _, v := range e.Entries {
		if v.SpentDate == date {
			// selection = append(selection, v)
			if IsWorkday(daySelector) {
				if v.Task.Id == 8814697 { // TODO: Switch to use variable from config.
					// saldoused = true
					saldo = saldo + v.Hours
				} else {
					hours = hours + v.Hours
				}

				// Calculate is the day full 7.5 hours, even if saldo using saldos.
			} else {
				hours = v.Hours
				overtime = v.Hours
			}
		}
	}
	if IsWorkday(daySelector) {
		if hours+saldo != 7.5 {
			overtime = (hours + saldo) - 7.5
		}
	}
	return
}

// DailyHours counts total logged in hours for selected date.
// Needs daySelector(time.Time) as parameter for selected date.
func (e *TimeEntries) DailyHours(daySelector time.Time) float64 {
	var selection Entries

	date := daySelector.Format("2006-01-02")

	for _, v := range e.Entries {
		if v.SpentDate == date {
			selection = append(selection, v)
		}
	}

	dayHours := selection.dayTotal()

	return dayHours
}

func (e Entries) dayTotal() float64 {
	var hours float64
	for _, v := range e {
		hours = hours + v.Hours
	}
	return hours
}

// IsWorkday functions as a helper function, to determine if selected date is workday or not.
func IsWorkday(date time.Time) bool { // Should this be placed in helpers.go?
	if date.Weekday().String() != "Saturday" && date.Weekday().String() != "Sunday" {
		return true
	}

	return false
}

// Filter is generic function to filter Entries
func (e *TimeEntries) Filter(f func(structs.Entries) bool) (ret []structs.Entries) {
	// var r []structs.Entries
	for _, v := range e.Entries {
		if f(v) == true {
			ret = append(ret, v)
		}
	}
	return
}

// func (e *TimeEntries) Filter(f func(structs.Entries) bool) []structs.Entries {
// 	var r Entries
// 	for _, v := range e.Entries {
// 		if f(v) == true {
// 			r = append(r, v)
// 		}
// 	}
// 	return r
// }

// func filter(s []student, f func(student) bool) []student {
// 	var r []student
// 	for _, v := range s {
// 		if f(v) == true {
// 			r = append(r, v)
// 		}
// 	}
// 	return r
// }

// func (t *TimeEntries) Choose(selector string, test func(string) bool) (ret []structs.Entries) {
// 	for _, v := range t.Entries {
// 		if test(s) {
// 			ret = append(ret, v)
// 		}
// 	}
// 	return
// }

// func (t *TimeEntries) Choose(s string, test func(string) bool) (ret []structs.Entries) {
// 	for _, v := range t.Entries {
// 		if test(s) {
// 			ret = append(ret, v)
// 		}
// 	}
// 	return
// }
