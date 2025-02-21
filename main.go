package main

import (
	dc "example-go-component/internal/edgee/components/data-collection"

	"go.bytecodealliance.org/cm"
)

// you should not need to modify this file
// this is a wrapper around the actual implementation located in component.go

type Result = cm.Result[dc.EdgeeRequestShape, dc.EdgeeRequest, string]

func resultWrapper(request dc.EdgeeRequest) (result Result) {
	return cm.OK[Result](request)
}

func init() {
	dc.Exports.Page = func(e dc.Event, cred dc.Dict) Result {
		return resultWrapper(PageHandler(e, cred))
	}
	dc.Exports.Track = func(e dc.Event, cred dc.Dict) Result {
		return resultWrapper(TrackHandler(e, cred))
	}
	dc.Exports.User = func(e dc.Event, cred dc.Dict) Result {
		return resultWrapper(UserHandler(e, cred))
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
