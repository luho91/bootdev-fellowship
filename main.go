package main

import (
	"fmt"
	"os"
)

func main() {
	token, err := getToken()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Obtained token successfully, submitting lessons")

	apiUrl, err := getApiUrl()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = sendLessons(apiUrl, token)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	os.Exit(0)
}
