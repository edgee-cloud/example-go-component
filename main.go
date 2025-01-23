package main

import (
	dc "example-go-component/internal/edgee/protocols/data-collection"
)

func init() {
	dc.Exports.Page = PageImpl
	dc.Exports.Track = TrackImpl
	dc.Exports.User = UserImpl
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
