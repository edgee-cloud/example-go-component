package main

import (
	incominghandler "example-go-component/internal/wasi/http/incoming-handler"
	"example-go-component/internal/wasi/io/streams"
)

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

