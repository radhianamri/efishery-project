package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Store initializes echo-session's RedisStore struct

var conf config

// DB initializes *gorm.DB struct
var DB *gorm.DB

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
