package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/renato-macedo/superheroapi/handlers/utils"
	"github.com/renato-macedo/superheroapi/services"
)

// CharacterHandler accepts the requests
type CharacterHandler struct {
	Service *services.CharacterService
}

// NewCharacterHandler return a instance with the given CharacterService
func NewCharacterHandler(service *services.CharacterService) *CharacterHandler {
	return &CharacterHandler{
		Service: service,
	}
}

// CreateCharacter handler
func (h *CharacterHandler) CreateCharacter(c *fiber.Ctx) {
	body := &utils.CreateRequest{}
	// Parse body into struct

	if err := c.BodyParser(body); err != nil {
		c.Status(400).JSON(utils.BadRequest("Invalid request body"))
		return
	}

	if body.Name == "" {
		c.Status(400).JSON(utils.BadRequest("name must not be blank"))
		return
	}

	createdCharacters, err := h.Service.Create(body.Name)

	if err != nil {
		c.Status(400).JSON(utils.BadRequest(err.Error()))
		return
	}
	if len(createdCharacters) == 0 {
		c.Status(400).JSON(utils.BadRequest("Character already exists"))
		return
	}
	response := utils.NewCreatedResponse(createdCharacters)

	if err := c.Status(201).JSON(response); err != nil {
		c.Status(500).Send(err)
		return
	}
}

// GetCharacter handles GET requests on /super
func (h *CharacterHandler) GetCharacter(c *fiber.Ctx) {
	characters := h.Service.FindAll()
	if err := c.Status(200).JSON(characters); err != nil {
		c.Status(500).Send(err)
		return
	}
}

// GetHeros handles GET requests on /super/heros
func (h *CharacterHandler) GetHeros(c *fiber.Ctx) {
	characters := h.Service.FindHeros()
	if err := c.Status(200).JSON(characters); err != nil {
		c.Status(500).Send(err)
		return
	}
}

// GetVillains handles GET requests on /super/villains
func (h *CharacterHandler) GetVillains(c *fiber.Ctx) {
	characters := h.Service.FindVillains()
	if err := c.Status(200).JSON(characters); err != nil {
		c.Status(500).Send(err)
	}

}
