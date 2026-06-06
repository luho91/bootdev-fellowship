package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
	"gopkg.in/yaml.v3"
)

func getToken() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please pick a method. Available methods:\n - Read from Bootdev CLI config file (1)\n - Paste token manually (2)\n - Use Credentials to obtain token (3)")
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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("failed to get home dir")
	}

	confPath := fmt.Sprintf("%s/.bootdev.yaml", homeDir)
	bootdevConfData := BootdevConfig{}

	data, err := os.ReadFile(confPath)
	if err != nil {
		return "", errors.New("failed to read config file")
	}

	err = yaml.Unmarshal(data, &bootdevConfData)
	if err != nil {
		return "", errors.New("failed to unmarshal config data")
	}

	if bootdevConfData.AccessToken == "" {
		return "", errors.New("failed to get token from config data")
	}

	return bootdevConfData.AccessToken, nil
}

func undefinedTokenMethod() (string, error) {
	return "", errors.New("undefined token method")
}
