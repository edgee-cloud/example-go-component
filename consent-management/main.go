package main

import (
	cmp "example-go-component/internal/edgee/components/consent-management"

	"go.bytecodealliance.org/cm"
)

// you should not need to modify this file
// this is a wrapper around the actual implementation located in component.go

type Option = cm.Option[cmp.Consent]

func None() Option {
	return cm.None[cmp.Consent]()
}

func Granted() Option {
	return cm.Some(cmp.ConsentGranted)
}

func Denied() Option {
	return cm.Some(cmp.ConsentDenied)
}

func Pending() Option {
	return cm.Some(cmp.ConsentPending)
}

func init() {
	cmp.Exports.Map = MapHandler
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
