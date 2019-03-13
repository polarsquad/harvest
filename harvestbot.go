package main

import (
	"fmt"

	// "net/http"
	// "encoding/json"

	"./harvest"
)

// There needs to moved to config.json file ASAP
var authToken = ""
var accountId = ""

func main() {
	h := harvest.Init(authToken, accountId)
	// h.Api.AuthToken = authToken
	// h.Api.AccountId = accountId

	user, err := h.GetUser()
	if err != nil {
		fmt.Printf("GetUser ERROR: %v", err)
	}
	h.User = user

	fmt.Printf("Name: %v\n", h.User.FirstName+" "+h.User.LastName)
	fmt.Printf("Email: %v\n", h.User.Email)
}
