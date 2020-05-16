package utils

import (
	"strings"

	"github.com/renato-macedo/superheroapi/domain"
)

// CharacterResponse character DTO
type CharacterResponse struct {
	*domain.Character
	Groups []string `json:"groups,omitempty"`
}

// CreatedResponse to be returned as json
type CreatedResponse struct {
	CharactersCount int                  `json:"characters_added"`
	Characters      []*CharacterResponse `json:"characters"`
}

// NewCreatedResponse takes the given slice an return a CreatedResponse
func NewCreatedResponse(characters []*domain.Character) *CreatedResponse {
	var supers []*CharacterResponse

	for _, value := range characters {
		supers = append(supers, &CharacterResponse{
			Character: value,
			Groups:    parseGroups(value.GroupAffiliation),
		})
	}
	return &CreatedResponse{
		CharactersCount: len(supers),
		Characters:      supers,
	}
}

func parseGroups(groupField string) []string {
	if groupField == "-" {
		return make([]string, 0)
	}
	// TODO handle edge cases
	return strings.Split(groupField, ", ")
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
