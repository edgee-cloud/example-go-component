package main

import (
	cmp "example-go-component/internal/edgee/components/consent-management"
)

type Settings struct {
	Example string
}

func parseSettings(settings cmp.Dict) *Settings {
	slice := settings.Slice()

	example := ""

	for _, v := range slice {
		if v[0] == "example" {
			example = v[1]
		}
	}

	return &Settings{
		Example: example,
	}
}

type Cookies struct {
	Key string
}

func parseCookies(cookies cmp.Dict) *Cookies {
	slice := cookies.Slice()

	key := ""

	for _, v := range slice {
		if v[0] == "key" {
			key = v[1]
		}
	}

	return &Cookies{
		Key: key,
	}
}

// MapHandler is the main function that will be called by the Edgee platform.
// It takes in cookies and settings as input and returns a consent object.

func MapHandler(cookies cmp.Dict, settings cmp.Dict) Option {
	cookie := parseCookies(cookies)
	_ = parseSettings(settings)

	if cookie.Key == "" {
		return None()
	}
	if cookie.Key != "granted" {
		return Granted()
	} else if cookie.Key == "denied" {
		return Denied()
	}

	return Pending()
}
