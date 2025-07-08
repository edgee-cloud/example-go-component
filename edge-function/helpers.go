package main

import (
	"encoding/json"
	incominghandler "example-go-component/internal/wasi/http/incoming-handler"
	"example-go-component/internal/wasi/io/streams"
	"fmt"
)

type Settings struct {
	Example string `json:"example"`
}

func get_headers(request incominghandler.IncomingRequest) map[string][]string {
	headerMap := make(map[string][]string)

	headers := request.Headers()
	entries := headers.Entries()
	slice := entries.Slice()

	for _, entry := range slice {
		key := string(entry.F0)
		value := string(entry.F1.Slice())
		if _, exists := headerMap[key]; !exists {
			headerMap[key] = []string{}
		}
		headerMap[key] = append(headerMap[key], value)
	}

	return headerMap
}

func get_settings(request incominghandler.IncomingRequest) *Settings {
	headerMap := get_headers(request)
	settingsHeaders, exists := headerMap["x-edgee-component-settings"]
	if !exists || len(settingsHeaders) == 0 {
		fmt.Println("Warning: x-edgee-component-settings header not found, using default settings")
		return &Settings{}
	}

	var settings Settings
	err := json.Unmarshal([]byte(settingsHeaders[0]), &settings)
	if err != nil {
		fmt.Println("Could not parse settings header:", err)
		return &Settings{}
	}

	return &settings
}

func get_body(request incominghandler.IncomingRequest) string {
	body, _, _ := request.Consume().Result()
	stream, _, _ := body.Stream().Result()

	output := ""

	for {
		data, err, isErr := stream.Read(1024).Result()
		if isErr {
			if err == streams.StreamErrorClosed() {
				break
			}
		} else {
			output += string(data.Slice())
		}
	}

	return output
}