package main

import (
	incominghandler "example-go-component/internal/wasi/http/incoming-handler"
	wasihttp "example-go-component/internal/wasi/http/types"
	"fmt"

	"encoding/json"
	"go.bytecodealliance.org/cm"
)

type Settings struct {
	Example string `json:"example"`
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
