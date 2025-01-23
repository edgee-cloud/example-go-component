package main

import (
	dc "example-go-component/internal/edgee/protocols/data-collection"
	"go.bytecodealliance.org/cm"
)

type Result = cm.Result[dc.EdgeeRequestShape, dc.EdgeeRequest, string]

func resultWrapper(request dc.EdgeeRequest) (result Result) {
	return cm.OK[Result](request)
}

func init() {
	dc.Exports.Page = func(e dc.Event, cred dc.Dict) Result {
		return resultWrapper(PageImpl(e, cred))
	}
	dc.Exports.Track = func(e dc.Event, cred dc.Dict) Result {
		return resultWrapper(TrackImpl(e, cred))
	}
	dc.Exports.User = func(e dc.Event, cred dc.Dict) Result {
		return resultWrapper(UserImpl(e, cred))
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
