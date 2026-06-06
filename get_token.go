package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
)

func getToken() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please pick a method (type only the number). Available methods:\n - Read from Bootdev CLI config file (1)\n - Paste token manually (2)\n - Use Credentials to obtain token (3)")
	fmt.Print("Your choice: ")
	scanner.Scan()
	input := scanner.Text()

	tokenFunc := undefinedTokenMethod

	switch input {
	case "1": tokenFunc = getTokenFromCLIConfig
	}

	token, err := tokenFunc()

	return token, err
}

func getTokenFromCLIConfig() (string, error) {
	cfg, err := readBootdevConfig()

	if cfg.AccessToken == "" {
		return "", errors.New("failed to get token from config data")
	}

	return cfg.AccessToken, err
}

func undefinedTokenMethod() (string, error) {
	return "", errors.New("undefined token method")
}
