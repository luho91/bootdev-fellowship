package main

import (
	"errors"
)

func getApiUrl() (string, error) {
	cfg, err := readBootdevConfig()

	if cfg.ApiUrl == "" {
		return "", errors.New("failed to get api url from config data")
	}

	return cfg.ApiUrl, err
}
