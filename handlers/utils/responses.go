package utils

import (
	"github.com/renato-macedo/superheroapi/domain"
	"strings"
)

type CharacterResponse struct {
	*domain.Character
	Groups []string `json:"groups, omitempty"`
}

type CreatedResponse struct {
	CharactersCount int                  `json:"characters_added"`
	Characters      []*CharacterResponse `json:"characters"`
}

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
