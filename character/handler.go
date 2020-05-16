package character

import (
	"strings"

	"github.com/gofiber/fiber"
	"github.com/renato-macedo/superheroapi/utils"
)

// CreateRequest dto
type CreateRequest struct {
	Name string `json:"name"`
}

// Handler accepts the requests
type Handler struct {
	Service *Service
}

// NewHandler return a instance with the given CharacterService
func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// CreateCharacter handler
func (h *Handler) CreateCharacter(c *fiber.Ctx) {
	body := &CreateRequest{}

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
		if err == utils.ErrSomethingWrong {
			c.Status(500).JSON(utils.ServerError("Something went wrong"))
			return
		}
		c.Status(400).JSON(utils.BadRequest(err.Error()))
		return
	}

	if len(createdCharacters) == 0 {
		c.Status(400).JSON(utils.BadRequest("Character already exists"))
		return
	}
	response := NewCreatedResponse(createdCharacters)

	if err := c.Status(201).JSON(response); err != nil {
		c.Status(500).Send(err)
		return
	}
}

// GetCharacter handles GET requests on /super
func (h *Handler) GetCharacter(c *fiber.Ctx) {
	characters := h.Service.FindAll()

	supers := NewSliceDTO(characters)

	if err := c.Status(200).JSON(supers); err != nil {
		c.Status(500).Send(err)
	}

	if err := c.Status(200).JSON(characters); err != nil {
		c.Status(500).Send(err)
		return
	}
}

// GetHeros handles GET requests on /super/heros
func (h *Handler) GetHeros(c *fiber.Ctx) {
	characters := h.Service.FindHeros()

	supers := NewSliceDTO(characters)

	if err := c.Status(200).JSON(supers); err != nil {
		c.Status(500).Send(err)
	}
}

// GetVillains handles GET requests on /super/villains
func (h *Handler) GetVillains(c *fiber.Ctx) {
	characters := h.Service.FindVillains()

	supers := NewSliceDTO(characters)

	if err := c.Status(200).JSON(supers); err != nil {
		c.Status(500).Send(err)
	}

}

//FindByID handles GET requests on /super/:id
func (h *Handler) FindByID(c *fiber.Ctx) {
	id := c.Params("id")
	character, err := h.Service.FindByID(id)

	if err != nil {
		c.Status(404).JSON(utils.NotFound("Character not found"))
		return
	}

	super := NewDTO(character)

	if err := c.JSON(super); err != nil {
		c.Status(500).JSON(utils.ServerError("Something went wrong"))
	}
}

// Search handles GET requests on /super/search
func (h *Handler) Search(c *fiber.Ctx) {
	name := c.Query("name")

	characters, err := h.Service.FindByName(strings.Title(strings.ToLower(name)))
	if err != nil {
		c.Status(404).JSON(utils.NotFound("Character not found"))
		return
	}

	supers := NewSliceDTO(characters)

	if err := c.JSON(supers); err != nil {
		c.Status(500).JSON(utils.ServerError("Something went wrong"))
	}
}

// Delete handle DELETE requests on /super/:id
func (h *Handler) Delete(c *fiber.Ctx) {
	id := c.Params("id")
	if id == "" {
		c.Status(400).JSON(utils.NotFound("Missing character id"))
		return
	}
	err := h.Service.Delete(id)
	if err != nil {
		c.Status(400).JSON(utils.NotFound("Character does not exists"))
		return
	}

	c.Status(200).JSON(Ok("Character deleted"))
}
