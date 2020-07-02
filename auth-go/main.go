package main

import (
	"log"

	"github.com/radhianamri/efishery-project/auth-go/config"
	"github.com/radhianamri/efishery-project/auth-go/db"
	"github.com/radhianamri/efishery-project/auth-go/server"
)

func main() {

	config.Init()
	log.Println(config.GetConfig())
	db.Init()
	server.Init()
}
