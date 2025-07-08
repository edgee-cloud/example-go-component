package main

import (
	"encoding/json"
	incominghandler "example-go-component/internal/wasi/http/incoming-handler"
	wasihttp "example-go-component/internal/wasi/http/types"
	"example-go-component/internal/wasi/io/streams"
	"fmt"

	"go.bytecodealliance.org/cm"
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

type Settings struct {
	Example string `json:"example"`
}

func get_settings(request incominghandler.IncomingRequest) *Settings {
	headerMap := get_headers(request)
	settingsHeader := headerMap["x-edgee-component-settings"][0]

	var settings Settings

	err := json.Unmarshal([]byte(settingsHeader), &settings)
	if err != nil {
		fmt.Println("Error:", err)
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

func Handle(request incominghandler.IncomingRequest, responseOut incominghandler.ResponseOutparam) {
	index := `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Coming Soon</title>
  <style>
    body {
      margin: 0;
      padding: 0;
      font-family: system-ui, sans-serif;
      display: flex;
      justify-content: center;
      align-items: center;
      text-align: center;
      height: 100vh;
      background-color: #f4f4f4;
      color: #333;
    }
    .container {
      max-width: 400px;
      padding: 2rem;
    }
    h1 {
      font-size: 2.5rem;
      margin-bottom: 1rem;
    }
    p {
      font-size: 1.1rem;
      margin-bottom: 2rem;
    }
    footer {
      font-size: 0.9rem;
      color: #888;
    }
    a {
      color: #007bff;
      text-decoration: none;
    }
    a:hover {
      text-decoration: underline;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>Coming Soon</h1>
    <p>We're working hard to launch something awesome. Stay tuned!</p>
    <footer>Served by <a href="https://www.edgee.cloud">Edgee</a></footer>
  </div>
</body>
</html>
`
	// get the headers, settings, and body from the request
	_ = get_headers(request)
	settings := get_settings(request)
	fmt.Println("Settings:", settings)
	incoming_body := get_body(request)
	fmt.Println("Request body:", incoming_body)

	// write the response
	bytes := []uint8(index)
	response := wasihttp.NewOutgoingResponse(wasihttp.NewFields())
	response.SetStatusCode(200)
	body, _, _ := response.Body().Result()
	stream, _, _ := body.Write().Result()

	wasihttp.ResponseOutparamSet(responseOut, cm.OK[cm.Result[wasihttp.ErrorCodeShape, wasihttp.OutgoingResponse, wasihttp.ErrorCode]](response))

	index2 := cm.NewList(&bytes[0], len(bytes))
	stream.Write(index2)
	stream.BlockingFlush()
	stream.ResourceDrop()
	_, _, iserr := wasihttp.OutgoingBodyFinish(body, cm.None[wasihttp.Trailers]()).Result()
	if iserr {
		panic("Failed to finish outgoing body")
	}
	println("Response finished successfully")
}
