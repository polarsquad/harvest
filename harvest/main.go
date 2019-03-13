package harvest

import (
	// "fmt"

	"./structs"
)

type Harvest struct {
	API     *structs.API
	User    *structs.User
	Project string
	Entries *[]structs.TimeEntries
}

// Init methot initializes the data structure needed for Harvest
func Init(authToken string, accountID string) *Harvest {
	a := &structs.API{
		AuthToken: authToken,
		AccountID: accountID,
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
