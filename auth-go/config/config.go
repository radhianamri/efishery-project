package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

var conf config

// DB initializes *gorm.DB struct
var DB *gorm.DB

const (
	//RegexOfficePhoneNumber example: +6209382929
	RegexOfficePhoneNumber = "^[+][0-9]{7,15}$"
)

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
