package structs

import "time"

// API struct has needed info for API authentication and the API URL
type API struct {
	AuthToken string
	AccountID string
	BaseURL   string
}

// Env is...
type Env struct {
	DateFormatter string
	FlexitimeIDs  []int64
}

// UserList is list of all users
type UserList struct {
	Users []User `json:"users"`
}

// User is a object of a user fetched from REST API
type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  bool   `json:"is_active"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt string `json:"created_at"`
}

// List of entries fetched from API
type TimeEntries struct {
	TotalEntries int64     `json:"total_entries"`
	TotalPages   int64     `json:"total_pages"`
	Links        Links     `json:"links"`
	Entries      []Entries `json:"time_entries"`
	DailyTotal   []dailyHours
}

// Time Entries paging links
type Links struct {
	First    string `json:"first"`
	Last     string `json:"last"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// type Entries map[string]interface{}
type Entries struct {
	Billable  bool           `json:"billable"`
	Client    EntriesClient  `json:"client"`
	Updated   string         `json:"updated_at"`
	Hours     float64        `json:"hours"`
	Id        int64          `json:"id"`
	IsBilled  bool           `json:"is_billed"`
	IsClosed  bool           `json:"is_closed"`
	Notes     string         `json:"notes"`
	Project   EntriesProject `json:"project"`
	SpentDate string         `json:"spent_date"`
	Task      EntriesTask    `json:"task"`
	User      EntriesUser    `json:"user"`
}

type EntriesClient struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type EntriesProject struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type EntriesTask struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type EntriesUser struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type dailyHours struct {
	Date       time.Time
	DailyHours float64
}

type selectedEntries struct {
	selection []Entries
}

type userAccount struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  bool   `json:"is_active"`
	IsAdmin   bool   `json:"is_admin"`
}
