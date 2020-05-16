package character

import (
	"strings"
)

// DTO character DTO
type DTO struct {
	Character
	Groups []string `json:"groups,omitempty"`
}

// CreatedResponse to be returned as json
type CreatedResponse struct {
	CharactersCount int   `json:"characters_added"`
	Characters      []DTO `json:"characters"`
}

// NewDTO creates a DTO with the given character model
func NewDTO(character *Character) DTO {
	return DTO{
		Character: *character,
		Groups:    parseGroups(character.GroupAffiliation),
	}
}

// NewSliceDTO transforms the slice of Character model into a slice of Character response
func NewSliceDTO(characters []Character) []DTO {
	var supers []DTO
	for _, value := range characters {
		supers = append(supers, NewDTO(&value))
	}

	return supers
}

// NewCreatedResponse takes the given slice an return a CreatedResponse
func NewCreatedResponse(characters []Character) *CreatedResponse {

	supers := NewSliceDTO(characters)
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

	groups := make([]string, 0)

	// two types of separators ', ' and '; '
	for _, firstLv := range strings.Split(groupField, ", ") {
		for _, secondLv := range strings.Split(firstLv, "; ") {
			groups = append(groups, secondLv)
		}
	}

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
