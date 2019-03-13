package main

import (
	"fmt"

	// "net/http"
	// "encoding/json"

	"./harvest"
	"./harvest/config"
)

func main() {
	configFile := "config.json"
	c := config.LoadConfig(configFile)

	h := harvest.Init(c)

	user, err := h.GetUser()
	if err != nil {
		fmt.Printf("GetUser ERROR: %v", err)
	}
	h.User = user

	fmt.Printf("Name: %v\n", h.User.FirstName+" "+h.User.LastName)
	fmt.Printf("Email: %v\n", h.User.Email)
}
