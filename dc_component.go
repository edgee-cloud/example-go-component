package main

import (
	dc "example-go-component/internal/edgee/protocols/data-collection"
	"go.bytecodealliance.org/cm"
)

// Implement the datacollection.Exports.Page, datacollection.Exports.Track, and datacollection.Exports.User functions.
// These functions are called by the Edgee runtime to get the HTTP request to make to the provider's API for each event type.
func PageHandler(e dc.Event, cred dc.Dict) dc.EdgeeRequest {
	// Access creds by using the Slice method
	// For example, if you component is setup as following:
	// [[components.data_collection]]
	// name = "my_component"
	// component = "outpout.wasm"
	// credentials.test_project_id = "123456789"
	// credentials.test_write_key = "abcdefg"
	// Then
	// cred.Slice() will return a slice of tuples with the following values:
	// [["test_project_id", "123456789"], ["test_write_key", "abcdefg"]]
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"Authorization", "Bearer token123"},
	}
	list := cm.NewList(&headers[0], len(headers))
	dict := dc.Dict(list)
	edgeeRequest := dc.EdgeeRequest{
		Method:  dc.HTTPMethodGET,
		URL:     "https://example.com/api/resource",
		Headers: dict,
		Body:    `{"key": "value"}`,
	}

	return edgeeRequest
}

func TrackHandler(e dc.Event, cred dc.Dict) dc.EdgeeRequest {
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"Authorization", "Bearer token123"},
	}
	list := cm.NewList(&headers[0], len(headers))
	dict := dc.Dict(list)
	edgeeRequest := dc.EdgeeRequest{
		Method:  dc.HTTPMethodGET,
		URL:     "https://example.com/api/resource",
		Headers: dict,
		Body:    `{"key": "value"}`,
	}

	return edgeeRequest
}

func UserHandler(e dc.Event, cred dc.Dict) dc.EdgeeRequest {
	headers := [][2]string{
		{"Content-Type", "application/json"},
		{"Authorization", "Bearer token123"},
	}
	list := cm.NewList(&headers[0], len(headers))
	dict := dc.Dict(list)
	edgeeRequest := dc.EdgeeRequest{
		Method:  dc.HTTPMethodGET,
		URL:     "https://example.com/api/resource",
		Headers: dict,
		Body:    `{"key": "value"}`,
	}

	return edgeeRequest
}
