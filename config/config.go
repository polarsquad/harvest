package config

import (
	"log"
	"os"

	"encoding/json"
)

// Config Has mainly API configurations, and maybe some global config values
type Config struct {
	API API `json:"api"`
	Env Env `json:"env"`
}

// API struct holds authentication information for Harvest REST API,
// and URL of the API
type API struct {
	AuthToken string `json:"auth_token"`
	AccountID string `json:"account_id"`
	BaseURL   string `json:"base_url"`
}

// Env has generic variables for configuration, eg. date formatter string, etc...
type Env struct {
	DateFormatter string `json:"date_formatter"`
}

// LoadConfig loads the file and parses it to struct
func LoadConfig(file string) *Config {
	c := new(Config)

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Printf("OPEN FILE ERROR: %v\n", err.Error())
		return c
	}

	confJSONParser := json.NewDecoder(configFile)
	confJSONParser.Decode(&c)

	return c
}
