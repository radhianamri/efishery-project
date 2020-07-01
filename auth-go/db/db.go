package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // to establish MySQL connection
	"github.com/jinzhu/gorm"
	"github.com/radhianamri/efishery-project/auth-go/config"
	"github.com/radhianamri/efishery-project/auth-go/model"
)

// DropTables -
func DropTables(db *gorm.DB) {
	config.DB.DropTableIfExists(
		&model.User{},
	)
}

// MigrateTables -
func MigrateTables(db *gorm.DB) {
	// Migrate
	config.DB.AutoMigrate(
		&model.User{},
	)

	config.DB.Model(&model.User{}).AddIndex(
		"idx_phone_pass", "phone", "password",
	)

}

// Init initialize DB connection
func Init() {
	var err error
	// Get config
	c := config.GetConfig()

	// Initializes MySQL database connection
	path := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=UTC",
		c.MySQLUser,
		c.MySQLPass,
		c.MySQLHost,
		c.MySQLPort,
	)
	config.DB, err = gorm.Open("mysql", path)
	if err != nil {
		log.Fatalf("Failed to establish connection to MySQL database: %s", err.Error())
	}
	// Set UTC timestamp
	config.DB.Exec("SET @@global.time_zone='+00:00';")
	config.DB.Exec("SET @@session.time_zone='+00:00';")

	// Initialize Schema
	config.DB.Exec("CREATE SCHEMA IF NOT EXISTS efishery")

	// if c.Mode == "test" {
	// 	DropTables(config.DB)
	// }

	MigrateTables(config.DB)
}
