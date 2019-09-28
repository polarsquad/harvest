package harvest

import (
	// "log"
	"time"
)

// TotalHours counts total logged hours from the TimeEntries struct
func (h *Harvest) TotalHours() float64 {
	var hours float64
	for _, v := range h.TimeEntries.Entries {
		hours = hours + v.Hours
	}
	// fmt.Printf("Total hours: %v\n", hours)
	return hours
}

// GetOvertime counts overtime hours from TimeEntries, using also dayTotal function as a helper.
func (e *TimeEntries) GetOvertime(from time.Time, to time.Time) {

	return
}

// DailyHours counts total logged in hours for selected date.
// Needs daySelector(time.Time) as parameter for selected day.
func (e *TimeEntries) DailyHours(daySelector time.Time) float64 {
	var selection Entries

	day, _ := time.Parse("2006-01-02", daySelector.String())

	for i := 0; i < len(e.Entries); i++ {
		if e.Entries[i].SpentDate == day.String() {
			selection = append(selection, e.Entries[i])
		}
	}

	for _, v := range e.Entries {
		if v.SpentDate == day.String() {
			selection = append(selection, v)
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

// isWorkday functions as a helper function, to determine if selected day is workday or not.
func isWorkday(date time.Time) bool { // Should this be placed in helpers.go?
	if date.Weekday().String() != "Saturday" && date.Weekday().String() != "Sunday" {
		return true
	}

	return false
}

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
