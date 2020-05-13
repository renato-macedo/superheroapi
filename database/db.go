package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/renato-macedo/superheroapi/domain"
	"log"
)

// Connect to the database
func Connect(dsn string) *gorm.DB {

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db.AutoMigrate(&domain.Character{})

	return db
}
