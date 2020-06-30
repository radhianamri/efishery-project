package server

import "github.com/radhianamri/efishery-project/auth-go/config"

// Init initialized config and and run the router accordingly
func Init() error {
	c := config.GetConfig()
	r := NewRouter()
	return r.Start(c.APIRouteURL)
}
