package main

import (
	"log"

	"github.com/radhianamri/efishery-project/auth-go/config"
	"github.com/radhianamri/efishery-project/auth-go/db"
	"github.com/radhianamri/efishery-project/auth-go/server"
)

// @title auth API
// @version 1.0
// @description This is an API documentation for auth

// @contact.email radhian.amri@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	config.Init()
	log.Println(config.GetConfig())
	db.Init()
	server.Init()
}
