package config

import (
	"fmt"
	"os"

	"encoding/json"
)

// Config Has mainly API configurations, and maybe some global config values
type Config struct {
	API API `json:"api"`
}

// API struct holds authentication information for Harvest REST API,
// and URL of the API
type API struct {
	AuthToken string `json:"auth_token"`
	AccountID string `json:"account_id"`
	BaseURL   string `json:"base_url"`
}

// LoadConfig loads the file and parses it to struct
func LoadConfig(file string) *Config {
	c := new(Config)

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Printf("OPEN FILE ERROR: %v\n", err.Error())
		return c
	}

	confJSONParser := json.NewDecoder(configFile)
	confJSONParser.Decode(&c)

	return c
}
