package character

import (
	"fmt"
	"log"
	"strings"

	superhero "github.com/renato-macedo/superheroapi/superhero"
	"github.com/renato-macedo/superheroapi/utils"
	uuid "github.com/satori/go.uuid"
)

// Service has methods that handles the Characters use cases
type Service struct {
	Repository Repository
	API        *superhero.Service
}

// Create method will store any character that matches the given name in tha superhero api
func (service *Service) Create(name string) ([]Character, error) {

	/*
		since it's possible to delete the super,
		it will be always necessary to check the super hero api, before adding new ones
	*/
	result, err := service.API.SearchCharacter(name)
	if err != nil {
		return nil, utils.ErrSomethingWrong
	}

	if result.Response == "error" {
		return nil, fmt.Errorf("could not add this super because it does not exist")
	}

	createdCharacters := make([]Character, 0)

	// check which of the search results were not stored yet
	for _, value := range result.Results {
		exists := service.Repository.HasCharacterWhere("api_id = ?", value.ID)

		if exists {
			continue
		}

		// if they were not stored just do it and append to the createdCharacters slice
		super := Character{
			ID:                uuid.NewV4().String(),
			APIID:             value.ID,
			Name:              value.Name,
			NameLowerCase:     strings.ToLower(value.Name),
			FullName:          value.Biography.FullName,
			Intelligence:      value.Powerstats.Intelligence,
			Power:             value.Powerstats.Power,
			Occupation:        value.Work.Occupation,
			Image:             value.Image.URL,
			Alignment:         value.Biography.Alignment,
			GroupAffiliation:  value.Connections.GroupAffiliation,
			NumberOfRelatives: parseRelatives(value.Connections.Relatives),
		}

		err := service.Repository.Store(&super)
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
func (service *Service) FindAll() []Character {
	return service.Repository.FindAll()
}

// FindHeros returns characters with alignment "good"
func (service *Service) FindHeros() []Character {
	return service.Repository.FindByFilter("alignment = ?", "good")
}

// FindVillains returns characters with alignment "bad"
func (service *Service) FindVillains() []Character {
	return service.Repository.FindByFilter("alignment = ?", "bad")
}

// FindByID return a character with the given id
func (service *Service) FindByID(id string) (*Character, error) {
	return service.Repository.FindByID(id)
}

// FindByName returns all characters that has the given name
func (service *Service) FindByName(name string) ([]Character, error) {
	return service.Repository.FindByName(name)
}

// Delete s a character from the database
func (service *Service) Delete(id string) error {
	return service.Repository.Delete(id)
}

func parseRelatives(relativesResult string) int {
	// TODO handle edge cases
	relatives := strings.Split(relativesResult, "),")
	return len(relatives)
}
