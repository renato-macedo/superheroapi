package main

import (
	"github.com/joho/godotenv"
	"github.com/renato-macedo/superheroapi/database"
	"github.com/renato-macedo/superheroapi/handlers"
	"github.com/renato-macedo/superheroapi/repos"
	"github.com/renato-macedo/superheroapi/services"
	"log"
	"os"

	"github.com/gofiber/fiber"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {

	DB := database.Connect(os.Getenv("ENV"))

	characterRepo := repos.NewCharacterRepo(DB)
	characterService := &services.CharacterService{
		Repository: characterRepo,
		API: &services.SuperHeroAPIService{
			BaseURL: "https://superheroapi.com/api",
			ApiKey:  os.Getenv("API_KEY"),
		},
	}

	characterHandler := handlers.NewCharacterHandler(characterService)

	app := fiber.New()

	app.Post("/super", characterHandler.CreateCharacter)

	app.Listen(3000)
}
