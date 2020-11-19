package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	PeakTemp    float64 `json:"peaktemp"`
	RegularTemp float64 `json:"regulartemp"`
}

// Retrieves a token from a local file.
func parseConfig(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	conf := &Config{}
	err = json.NewDecoder(f).Decode(conf)
	return conf, err
}
