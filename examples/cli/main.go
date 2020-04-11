package main

import (
	"fmt"
	"log"
	"time"

	"github.com/polarsquad/harvest"
	"github.com/polarsquad/harvest/config"

	// "github.com/ransoni/harvest/config"
	// "github.com/ransoni/harvest/config"

	"github.com/manifoldco/promptui"
)

var timeEntries harvest.TimeEntries
var h *harvest.Harvest
var user *harvest.User

// var c harvest.Config
var c *config.Config

var err error

func main() {

	configFile := "config.json"
	c = config.LoadConfig(configFile)

	c.API.AuthToken = "1990007.pt.XThqjXYrWtmn5vS1BCQ5HyvD7I9VLJh0T7CE8IWFo6_IpZ-oSCOgpya-skdOFcs1r0baqSLs4yEmvTqbJj2Log" //personal token
	// c.API.AuthToken = "2077799.pt.lD0F7ZgBPS4uI-lNNMCFzmw8-r7PNx-PeX3JtIC-Ld-ktAjvYQei4TZz_6azI6-gxrD0W_U6P_4jJwprOBCt1w" // Bot token
	c.API.AccountID = "851672"

	// h := harvest.InitHarvest(&c)
	h = harvest.Init(c)

	// user, err = h.GetUserByEmail("mika.tuominen@polarsquad.com")
	// if err != nil {
	// 	log.Fatalf("%s", err)
	// 	return
	// }

	user, err = h.GetUser()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Hello %s\n", user.FirstName)

	prompt := promptui.Select{
		Label: "Selec function",
		Items: []string{"Get User", "Get Flexitime"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Promp failed %v\n", err)
	}

	fmt.Printf("You choosed: %q\n", result)

	// if result == "Get Projects" {
	// 	// var projects *harvest.Projects
	// 	projects := h.GetProjects()

	// 	fmt.Printf("Number of projects: %v\n\n", len(projects.Projects))

	// 	for _, v := range projects.Projects {
	// 		fmt.Printf("Project name: %v\n\tClient: %v , ", v.Name, v.Client.Name)
	// 	}
	// 	// exit
	// }

	// if result == "Get Tasks" {
	// 	// var projects *harvest.Projects
	// 	tasks := h.GetTasks()

	// 	fmt.Printf("Number of tasks: %v\n\n", len(tasks.Tasks))
	// 	fmt.Printf("Flexitime ID should be: %v\n", h.Env.FlexitimeIDs)

	// 	for _, v := range tasks.Tasks {
	// 		if strings.Contains(v.Name, "aldovapaa") {
	// 			fmt.Printf("Task name: %v (ID: %v)\n", v.Name, v.Id)
	// 		}
	// 	}
	// 	// exit
	// }

	if result == "Get Flexitime" {
		getFlexitime()
	}

	// if result == "Legacy" {
	// 	legacy()
	// }
}

func getFlexitime() {
	fmt.Println(user)
	// createdAt, err := time.Parse(time.RFC3339, user.CreatedAt)
	// if err != nil {
	// 	fmt.Printf("[ERROR] %s", err)
	// }
	// // createdDate, err := createdAt.Format(c.Env.DateFormatter)
	// fmt.Printf("User: %s, Created: %s", user.FirstName, createdAt.Format(c.Env.DateFormatter))

	var startDate string
	fmt.Printf("Enter start date (eg. 2020-01-22): ")
	fmt.Scanf("%s", &startDate)
	var endDate string
	fmt.Printf("Enter end date (eg. 2020-01-22): ")
	fmt.Scanf("%s", &endDate)

	h.User, _ = h.GetUser()

	from, _ := time.Parse(c.Env.DateFormatter, startDate)
	to, err := time.Parse(c.Env.DateFormatter, endDate)
	// from, _ := time.Parse(c.Env.DateFormatter, "2020-01-01")
	// to, err := time.Parse(c.Env.DateFormatter, "2020-02-26")

	if err != nil {
		log.Println(err.Error)
	}

	// fmt.Printf("FROM: %s\n", from.Format(c.Env.DateFormatter))
	// fmt.Printf("TO: %s\n", to.Format(c.Env.DateFormatter))
	fmt.Printf("FROM: %s\n", from.Format("2006-01-02"))
	fmt.Printf("TO: %s\n", to.Format("2006-01-02"))

	// timeEntries := h.GetEntries(from, to, user)
	var usr = *user
	timeEntries := h.GetEntries(from, to, usr)

	overtime := timeEntries.GetOvertime(from, to)
	fmt.Printf("Total overtime is: %.1f\n", overtime)
}

// func legacy() {
// 	fmt.Println(user)
// 	createdAt, err := time.Parse(time.RFC3339, user.CreatedAt)
// 	if err != nil {
// 		fmt.Printf("[ERROR] %s", err)
// 	}
// 	// createdDate, err := createdAt.Format(c.Env.DateFormatter)
// 	fmt.Printf("User: %s, Created: %s", user.FirstName, createdAt.Format(c.Env.DateFormatter))

// 	var str string
// 	fmt.Scanf("%s", &str)

// 	h.User, _ = h.GetUser()

// 	from, _ := time.Parse(c.Env.DateFormatter, "2020-01-01")
// 	to, err := time.Parse(c.Env.DateFormatter, "2020-02-26")
// 	// from, _ := time.Parse("2006-01-02", "2019-06-04")
// 	// to, err := time.Parse("2006-01-02", "2019-12-31")
// 	if err != nil {
// 		log.Println(err.Error)
// 	}

// 	from = createdAt
// 	to = createdAt.AddDate(0, 3, 0)
// 	// fmt.Printf("FROM: %s\n", from.Format(c.Env.DateFormatter))
// 	// fmt.Printf("TO: %s\n", to.Format(c.Env.DateFormatter))
// 	fmt.Printf("FROM: %s\n", from.Format("2006-01-02"))
// 	fmt.Printf("TO: %s\n", to.Format("2006-01-02"))

// 	// timeEntries := h.GetEntries(from, to, user)
// 	timeEntries := h.GetEntries(from, to, user)

// 	fmt.Println("Sorting entries...")

// 	// Get first day when there are time entries
// 	var selectedDate time.Time
// 	fmt.Printf("\n*** Getting the first time entry... ***")

// 	for i := 0; i <= len(timeEntries.Entries); i++ {
// 		// selectedDate, err := time.Parse(c.Env.DateFormatter, timeEntries.Entries[i].SpentDate)
// 		// if err != nil {
// 		// 	log.Printf("[ERROR] Could not parse the time. %s", err)
// 		// }

// 		// if timeEntries.DailyHours(selectedDate) != 0 {
// 		if timeEntries.DailyHours(from.AddDate(0, 0, i)) != 0 {
// 			selectedDate = from.AddDate(0, 0, i)
// 			// selectedDate, err := time.Parse(c.Env.DateFormatter, timeEntries.Entries[i].SpentDate)
// 			// if err != nil {
// 			// 	log.Printf("[ERROR] Could not parse the time. %s", err)
// 			// }
// 			fmt.Printf("\nFirst day with entries is %s.\n", selectedDate.Format(c.Env.DateFormatter))
// 			fmt.Printf("Hours on that date: %v\n", timeEntries.DailyHours(selectedDate))
// 			// fmt.Printf("\nFirst day with entries is %s.\n", selectedDate.Format(c.Env.DateFormatter))
// 			break
// 		}
// 	}

// 	// fmt.Printf("Total Hours: %v\n", h.TotalHours())
// 	fmt.Printf("Total Hours: %v\n", timeEntries.TotalHours())
// 	fmt.Printf("Overtime (%v) between %s - %s: %v\n", h.Env.FlexitimeIDs, selectedDate.Format(c.Env.DateFormatter), to.Format(c.Env.DateFormatter), timeEntries.GetOvertime(selectedDate, to))

// 	// fmt.Printf("Num of Entries: %d", len(h.TimeEntries.Entries))
// 	fmt.Printf("Num of Entries: %d", len(timeEntries.Entries))

// 	// *** PRINT ALL ENTRIES ***
// 	//
// 	// var hours, saldo, overtime, totalOvertime, totalSaldo float64
// 	// for d := from; d.Before(to) || d.Equal(to); d = d.AddDate(0, 0, 1) {

// 	// 	// hours, saldo, overtime = h.TimeEntries.DailyTotals(d)
// 	// 	hours, saldo, overtime = timeEntries.DailyTotals(d)
// 	// 	totalOvertime = totalOvertime + overtime
// 	// 	totalSaldo = totalSaldo + saldo
// 	// 	fmt.Printf("\n%v %v: %v hours.", d.Weekday().String()[:3], d.Format("01.02.2006"), hours)
// 	// 	// fmt.Printf("\n%v %v: %v hours.", d.Weekday().String()[:3], d.Format(c.Env.DateFormatter), hours)
// 	// 	// if saldo != 0 {
// 	// 	fmt.Printf(", saldo: %v", saldo)
// 	// 	// }
// 	// 	// if hours+saldo != 7.5 {
// 	// 	fmt.Printf(" (%v)", overtime)
// 	// 	// printDay(h.TimeEntries.Entries, d)
// 	// 	printDay(timeEntries.Entries, d)

// 	// 	// }

// 	// 	// if harvest.IsWorkday(d) {
// 	// 	// 	if hours > 7.5 {
// 	// 	// 		fmt.Printf(" (+%v)", dailyHours-7.5)
// 	// 	// 		overtime = overtime + dailyHours - 7.5
// 	// 	// 	}
// 	// 	// 	if dailyHours < 7.5 {
// 	// 	// 		fmt.Printf(" (%v)", dailyHours-7.5)
// 	// 	// 		overtime = overtime + dailyHours - 7.5
// 	// 	// 	}
// 	// 	// }
// 	// }
// 	//
// 	// *** PRINT ALL ENTRIES END ***

// 	// for d := from; d.Before(to) || d.Equal(to); d = d.AddDate(0, 0, 1) {
// 	// 	dailyHours := h.TimeEntries.DailyHours(d)
// 	// 	fmt.Printf("\n%v %v: %v hours.", d.Weekday().String()[:3], d.Format(c.Env.DateFormatter), dailyHours)
// 	// 	if harvest.IsWorkday(d) {
// 	// 		if dailyHours > 7.5 {
// 	// 			fmt.Printf(" (+%v)", dailyHours-7.5)
// 	// 			overtime = overtime + dailyHours - 7.5
// 	// 		}
// 	// 		if dailyHours < 7.5 {
// 	// 			fmt.Printf(" (%v)", dailyHours-7.5)
// 	// 			overtime = overtime + dailyHours - 7.5
// 	// 		}
// 	// 	}
// 	// }

// 	// var totalOvertime, totalSaldo float64
// 	// fmt.Printf("\nTotalOvertime: %.1f", totalOvertime)
// 	// fmt.Printf("\nTotalSaldo: %.1f", totalSaldo)
// 	// fmt.Printf("\nTOTAL: %.1f", totalOvertime-totalSaldo)

// 	// fmt.Printf("\nHarvest.GetOvertime: %.1f", h.TimeEntries.GetOvertime(from, to))
// 	fmt.Printf("\nHarvest.GetOvertime: %.1f", timeEntries.GetOvertime(selectedDate, to))

// 	// mytest := func(s string) bool { return !strings.HasPrefix(s, "foo_") && len(s) <= 7 }
// 	// mytest := func(v structs.Entries) bool { return v.SpentDate == "2019-02-10" }

// 	// *** SOME SAMPLE TESTING
// 	// body, err := ioutil.ReadFile("sample/sample_timeEntries.json")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// 	// fmt.Print("ERROR: %v", err.Error)
// 	// }

// 	// var times harvest.TimeEntries
// 	// json.Unmarshal(body, &times)

// 	// fmt.Printf("\nNro of entries: %v", len(times.Entries))
// 	// *** SAMPLE TESTING END

// 	// for _, v := range times.Entries {
// 	// 	fmt.Printf("\n%s", v.SpentDate)
// 	// 	fmt.Printf("\n   Notes: %s", v.Notes)
// 	// 	selectedDay, _ = time.Parse(c.Env.DateFormatter, v.SpentDate)
// 	// 	fmt.Printf("\n   Hours: %v", times.DailyHours(selectedDay))
// 	// }

// 	// sel := h.TimeEntries.Filter(func(e structs.Entries) bool {
// 	// 	if e.SpentDate >= from.Format(c.Env.DateFormatter) && e.SpentDate <= from.AddDate(0, 0, 4).Format(c.Env.DateFormatter) {
// 	// 		return true
// 	// 	}
// 	// 	return false
// 	// })

// 	// sel := h.TimeEntries.Filter(dateSelector(from, to.AddDate(0, 0, 4)))
// 	sel := timeEntries.Filter(dateSelector(selectedDate, from.AddDate(0, 0, 5)))

// 	timeEntries.Entries = sel
// 	fmt.Printf("\n\nTotal(%v): %v", len(timeEntries.Entries), timeEntries.Total())

// 	// for _, v := range timeEntries.Entries {
// 	// 	fmt.Printf("\nDate: %v, Hours: %v", v.SpentDate, v.Hours)
// 	// }

// 	fmt.Printf("\nSelection total: %v", timeEntries.Total())

// 	// sel := times.TimeEntries.Filter(func(e structs.Entries) bool {
// 	// 	if e.Billable == true {
// 	// 		return true
// 	// 	}
// 	// 	return false
// 	// })
// 	// fmt.Println(sel)
// }

// func dateSelector(from time.Time, to time.Time) func(e structs.Entries) bool {
// 	return func(e structs.Entries) bool {
// 		if e.SpentDate >= from.Format("2006-01-02") && e.SpentDate <= to.Format("2006-01-02") {
// 			return true
// 		}
// 		return false
// 	}
// }

// func printDay(e []structs.Entries, day time.Time) {
// 	for _, v := range e {
// 		if v.SpentDate == day.Format("2006-01-02") {
// 			fmt.Printf("\n\t%v: %v (%s)", v.SpentDate, v.Hours, v.Task.Name)
// 		}
// 	}
// }
