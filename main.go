package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/renato-macedo/superheroapi/database"
	"github.com/renato-macedo/superheroapi/handlers"
	"github.com/renato-macedo/superheroapi/repos"
	"github.com/renato-macedo/superheroapi/services"

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
			APIKey:  os.Getenv("API_KEY"),
		},
	}

	characterHandler := handlers.NewCharacterHandler(characterService)

	app := fiber.New()

	app.Get("/super", characterHandler.GetCharacter)
	app.Get("/super/heros", characterHandler.GetHeros)
	app.Get("/super/villains", characterHandler.GetVillains)
	app.Get("/super/search", characterHandler.Search)
	app.Get("/super/:id", characterHandler.FindByID)

	app.Post("/super", characterHandler.CreateCharacter)
	app.Delete("/super/:id", characterHandler.Delete)

	app.Listen(3000)
}
