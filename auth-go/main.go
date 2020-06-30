package main

import (
	"github.com/radhianamri/efishery-project/auth-go/config"
	"github.com/radhianamri/efishery-project/auth-go/db"
	"github.com/radhianamri/efishery-project/auth-go/server"
)

func main() {
	config.Init()
	db.Init()
	server.Init()
}
