package services

import (
	"fmt"
	"log"
	"strings"

	"github.com/renato-macedo/superheroapi/domain"
	"github.com/renato-macedo/superheroapi/repos"
	uuid "github.com/satori/go.uuid"
)

// CharacterService has methods that handles the Characters use cases
type CharacterService struct {
	Repository repos.CharacterRepository
	API        *SuperHeroAPIService
}

// Create method will store any character that matches the given name in tha superhero api
func (service *CharacterService) Create(name string) ([]*domain.Character, error) {

	/*
		since it's possible to delete the super,
		it will be always necessary to check the super hero api, before adding new ones
	*/
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

		if exists {
			continue
		}

		// if they were not stored just do it and append to the createdCharacters slice
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

		if err != nil {
			log.Println(err)
			continue
		}
		createdCharacters = append(createdCharacters, super)

	}

	// return which character were added
	return createdCharacters, nil
}

// FindAll returns all stored supers
func (service *CharacterService) FindAll() []*domain.Character {
	return service.Repository.FindAll()
}

// FindHeros returns characters with alignment "good"
func (service *CharacterService) FindHeros() []*domain.Character {
	return service.Repository.FindByFilter("alignment = ?", "good")
}

// FindVillains returns characters with alignment "bad"
func (service *CharacterService) FindVillains() []*domain.Character {
	return service.Repository.FindByFilter("alignment = ?", "bad")
}

func parseRelatives(relativesResult string) int {
	// TODO handle edge cases
	relatives := strings.Split(relativesResult, "),")
	return len(relatives)
}
