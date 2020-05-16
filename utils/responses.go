package utils

import (
	"strings"

	"github.com/renato-macedo/superheroapi/domain"
)

// CharacterDTO character DTO
type CharacterDTO struct {
	domain.Character
	Groups []string `json:"groups,omitempty"`
}

// CreatedResponse to be returned as json
type CreatedResponse struct {
	CharactersCount int            `json:"characters_added"`
	Characters      []CharacterDTO `json:"characters"`
}

// NewCharacterDTO creates a CharacterDTO with the given character model
func NewCharacterDTO(character *domain.Character) CharacterDTO {
	return CharacterDTO{
		Character: *character,
		Groups:    parseGroups(character.GroupAffiliation),
	}
}

// NewSliceCharacterDTO transforms the slice of Character model into a slice of Character response
func NewSliceCharacterDTO(characters []domain.Character) []CharacterDTO {
	var supers []CharacterDTO
	for _, value := range characters {
		supers = append(supers, NewCharacterDTO(&value))
	}

	return supers
}

// NewCreatedResponse takes the given slice an return a CreatedResponse
func NewCreatedResponse(characters []domain.Character) *CreatedResponse {

	supers := NewSliceCharacterDTO(characters)
	return &CreatedResponse{
		CharactersCount: len(supers),
		Characters:      supers,
	}
}

// ParseGroups split the group affilliation into a slice
func parseGroups(groupField string) []string {
	if groupField == "-" {
		return make([]string, 0)
	}

	groups := strings.Split(groupField, ", ")

	// TODO handle edge cases
	return groups
}

// OkResponse the be returned on successful operations
type OkResponse struct {
	Message string `json:"message"`
}

// Ok returns a OkResponse with the given message
func Ok(message string) *OkResponse {
	return &OkResponse{
		Message: message,
	}
}
