package services

import (
	"fmt"
	"github.com/renato-macedo/superheroapi/domain"
	"github.com/renato-macedo/superheroapi/repos"
	"github.com/satori/go.uuid"
	"log"
	"strings"
)

type CharacterService struct {
	Repository repos.CharacterRepository
	API        *SuperHeroAPIService
}

func (service *CharacterService) Create(name string) ([]*domain.Character, error) {

	// it will be always necessary to check the super hero api
	result, err := service.API.SearchCharacter(name)
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	if result.Response == "error" {
		return nil, fmt.Errorf("could not add this super because it does not exist")
	}

	// var createdCharacters []*domain.Character

	createdCharacters := make([]*domain.Character, 0)

	// check which of the search results were not stored yet
	for _, value := range result.Results {
		exists := service.Repository.HasCharacterWhere("api_id = ?", value.ID)

		// if the were not stored just store them and append to the createdCharacters slice
		if !exists {
			super := &domain.Character{
				ID:                uuid.NewV4().String(),
				ApiId:             value.ID,
				Name:              value.Name,
				FullName:          value.Biography.FullName,
				Intelligence:      value.Powerstats.Intelligence,
				Power:             value.Powerstats.Power,
				Occupation:        value.Work.Occupation,
				Image:             value.Image.URL,
				Alignment:         value.Biography.Alignment,
				GroupAffiliation:  value.Connections.GroupAffiliation,
				NumberOfRelatives: parseRelatives(value.Connections.Relatives),
			}

			err := service.Repository.Store(super)
			log.Println(err)
			if err == nil {
				createdCharacters = append(createdCharacters, super)
			}
		}

	}

	// return which character were added
	return createdCharacters, nil
}



func parseRelatives(relativesResult string) int {
	// TODO handle edge cases
	relatives := strings.Split(relativesResult, "),")
	return len(relatives)
}
