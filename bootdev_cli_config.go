package main

import (
	"os"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
)

type BootdevConfig struct {
	AccessToken	string	`yaml:"access_token"`
	ApiUrl		string	`yaml:"api_url"`
}

func readBootdevConfig() (BootdevConfig, error) {
	cfg := BootdevConfig{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return cfg, errors.New("failed to get home dir")
	}

	confPath := fmt.Sprintf("%s/.bootdev.yaml", homeDir)

	data, err := os.ReadFile(confPath)
	if err != nil {
		return cfg, errors.New("failed to read config file")
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, errors.New("failed to unmarshal config data")
	}

	return cfg, nil
}
