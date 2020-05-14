package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/renato-macedo/superheroapi/domain"
	"log"
	"os"
)

// Connect to the database
func Connect(env string) *gorm.DB {

	var connection string
	if env != "test" {
		connection = os.Getenv("CONNECTION")
	} else {
		connection = os.Getenv("CONNECTION_TEST")
	}

	db, err := gorm.Open("postgres", connection)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db.AutoMigrate(&domain.Character{})

	return db
}
