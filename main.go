package harvest

import (
	// "fmt"

	"github.com/polarsquad/harvest/config"
	"github.com/polarsquad/harvest/structs"
)

// VERSION v0.0.2
const VERSION = "v0.0.2"

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
