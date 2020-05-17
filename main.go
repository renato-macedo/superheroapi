package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/renato-macedo/superheroapi/character"
	"github.com/renato-macedo/superheroapi/database"
	"github.com/renato-macedo/superheroapi/superhero"

	"github.com/gofiber/fiber"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	app := startApp()

	app.Listen(3000)
}

func startApp() *fiber.App {
	DB := database.Connect(os.Getenv("ENV"))

	characterRepo := character.NewCharacterRepo(DB)

	characterService := &character.Service{
		Repository: characterRepo,
		API: &superhero.Service{
			BaseURL: "https://superheroapi.com/api",
			APIKey:  os.Getenv("API_KEY"),
		},
	}

	characterHandler := character.NewHandler(characterService)

	app := fiber.New()

	app.Get("/super", characterHandler.GetCharacter)
	app.Get("/super/heros", characterHandler.GetHeros)
	app.Get("/super/villains", characterHandler.GetVillains)
	app.Get("/super/search", characterHandler.Search)
	app.Get("/super/:id", characterHandler.FindByID)

	app.Post("/super", characterHandler.CreateCharacter)
	app.Delete("/super/:id", characterHandler.Delete)

	return app
}
