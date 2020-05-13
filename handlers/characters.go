package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/renato-macedo/superheroapi/handlers/utils"
	"github.com/renato-macedo/superheroapi/services"
)

type CharacterHandler struct {
	Service *services.CharacterService
}

func NewCharacterHandler(service *services.CharacterService) *CharacterHandler {
	return &CharacterHandler{
		Service: service,
	}
}

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

	response := utils.NewCreatedResponse(createdCharacters)

	if err := c.Status(201).JSON(response); err != nil {
		c.Status(500).Send(err)
		return
	}
}