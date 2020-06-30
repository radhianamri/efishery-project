package config

import (
	"log"
)

// Store initializes echo-session's RedisStore struct

// DB initializes *gorm.DB struct

var conf config

// AWSClientSession initializes *session.Session struct

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() {
	if err := LoadTomlConfig(); err != nil {
		log.Fatal("Error loading toml config", err)
	}
}

func GetConfig() config {
	return conf
}
