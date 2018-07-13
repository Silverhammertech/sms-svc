package rest

import (
	oauth2 "github.com/Silverhammertech/oauth-lib"
)

var routes = ConfiguredRoutes{
	Route{
		Name:        "SMS",
		Method:      "POST",
		Pattern:     "/send",
		HandlerFunc: HandlePostMessage},
}

type ConfiguredRoutes []Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc oauth2.AuthHandler
}