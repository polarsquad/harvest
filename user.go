package harvest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "github.com/polarsquad/harvest/structs"
)

// GetUser Fetches the information of the logged in user
func (h *Harvest) GetUser() (*User, error) {

	url := "https://api.harvestapp.com/v2/users/me"
	// url := "https://api.harvestapp.com/v2/company"
	Client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Go Harvest API Sample")
	req.Header.Set("Harvest-Account-ID", h.API.AccountID)
	req.Header.Set("Authorization", "Bearer "+h.API.AuthToken)

	resp, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Couldn't fetch user API: %v", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("Can't read body: %v", err)
	// }

	defer resp.Body.Close()

	// body, err := getURL("GET", url)
	// if err != nil {
	// 	os.Exit(1)
	// }

	// var jsonResponse map[string]interface{}
	// var user structs.User
	var user User

	// json.Unmarshal(body, &jsonResponse)
	json.Unmarshal(body, &user)

	// prettyJson, _ := json.MarshalIndent(user, "", "  ")
	// fmt.Println(string(prettyJson))
	if !user.IsActive {
		log.Fatalf("User not active!")
	}

	return &user, nil
}

// GetUserByEmail is ...
func (h *Harvest) GetUserByEmail(email string) (*User, error) {
	var user User

	body, err := h.getURL("GET", usersURL)
	if err != nil {
		log.Fatalf("[ERROR] Could not get users.")
		return &user, err
	}

	var usersList Users
	json.Unmarshal(body, &usersList)

	for _, v := range usersList.Users {
		if v.Email == email {
			// user = v
			return &v, nil
		}
	}

	return &user, fmt.Errorf("could not found user")

}
