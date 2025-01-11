package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	StateDir  string `json:"dir"`
	StateFile string `json:"filename"`
}

func GettingConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := new(Config)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
