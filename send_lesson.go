package main

import (
	"encoding/json"
	"time"
	"net/url"
	"io"
	"net/http"
	"bytes"
	"fmt"
)

type submitData struct {
	lessonData
	CodeLastModAt	string	`json:"codeLastModAt"`
	UserTimezone	string	`json:"userTimezone"`
}

type lessonData struct {
	Input	string		`json:"input"`
	Files	[]fileData	`json:"files"`
}

type fileData struct {
	Name		string	`json:"Name"`
	Content		string	`json:"Content"`
	IsHidden	bool	`json:"IsHidden"`
	IsReadonly	bool	`json:"IsReadonly"`
}

func sendLessons(apiUrl, apiToken string) error {
	payload := submitData{}

	payload.UserTimezone = "Europe/Berlin"

	tz, err := time.LoadLocation(payload.UserTimezone)
	if err != nil {
		return err
	}

	t := time.Now()
	
	payload.CodeLastModAt = t.In(tz).Format(time.RFC3339)

	for lessonID, lessonObject := range Lessons {
		requestURL, err := url.JoinPath(apiUrl, "v1/lessons", lessonID)
		if err != nil {
			return err
		}

		payload.Files = lessonObject.Files
		payload.Input = lessonObject.Input

		body, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(body))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(responseBody))
	}
	return nil
}

var Lessons = map[string]lessonData {
	"78b4646f-85aa-42c7-ba46-faec2f0902a9": {
		Input: "Welcome to Fantasy Quest!",
		Files: []fileData {
			{
				Name:		"main.py",
				Content:	"print(\"Welcome to Fantasy Quest\")",
				IsHidden:	false,
				IsReadonly:	false,
			},
		},
	},
}
