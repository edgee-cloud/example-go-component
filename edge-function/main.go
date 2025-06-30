package main

import (
	incominghandler "example-go-component/internal/wasi/http/incoming-handler"
)

// you should not need to modify this file
// this is a wrapper around the actual implementation located in component.go

func init() {
	// 	Handle func(request IncomingRequest, responseOut ResponseOutparam)

	incominghandler.Exports.Handle = func(request incominghandler.IncomingRequest, responseOut incominghandler.ResponseOutparam) {
		Handle(request, responseOut)
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
