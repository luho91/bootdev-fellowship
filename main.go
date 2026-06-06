package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi, friend")

	token, err := getToken()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Token: %s\n", token)
	os.Exit(0)
}
