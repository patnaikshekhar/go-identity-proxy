package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config describes the config object for the proxy
type Config struct {
	Address          string `json:"address"`
	JwksURL          string `json:"jwksURL"`
	ExpectedAudience string `json:"expectedAudience"`
}

// ReadConfig reads the config from config.json
func ReadConfig() (*Config, error) {

	var result *Config

	// Read the file
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf(
			"Could not read config.json file %s", err.Error())
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf(
			"Invalid format of config.json %s", err.Error())
	}

	return result, nil
}
