package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // make linter happy
	"github.com/renato-macedo/superheroapi/character"
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

	db.AutoMigrate(&character.Character{})

	return db
}
