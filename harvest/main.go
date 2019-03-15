package harvest

const VERSION = "V0.1"

import (
	// "fmt"

	"./config"
	"./structs"
)

// Harvest creates the struct for the API, User and Entries
type Harvest struct {
	API     *structs.API
	User    *structs.User
	Project string
	Entries *[]structs.TimeEntries
}

// Init methot initializes the data structure needed for Harvest
func Init(conf *config.Config) *Harvest {
	a := &structs.API{
		AuthToken: conf.API.AuthToken,
		AccountID: conf.API.AccountID,
	}

	e := &[]structs.TimeEntries{}

	h := &Harvest{
		API:     a,
		User:    &structs.User{},
		Project: "",
		Entries: e,
	}
	// h := &Harvest{
	// 	User:     "Mika",
	// 	Projects: "Client",
	// 	Entries:  []TimeEntries{},
	// }

	return h
}
