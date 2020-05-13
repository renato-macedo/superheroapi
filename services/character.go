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
	Repository *repos.CharacterRepositoryDB
	API        *SuperHeroAPIService
}

func (service *CharacterService) Create(name string) ([]*domain.Character, error) {
	// characters := service.Repository.FindByName(strings.Title(strings.ToLower(name)))

	//if len(characters) >= 1 {
	//	return characters, fmt.Errorf("character already exists")
	//}

	result, err := service.API.SearchCharacter(name)
	if err != nil || result.Response == "error" {
		return nil, fmt.Errorf("could not add this super because it does not exists")
	}

	var createdCharacters []*domain.Character

	// var waitGroup sync.WaitGroup

	for _, value := range result.Results {
		super := &domain.Character{
			ID:                uuid.NewV4().String(),
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
		// waitGroup.Add(1)

		go func() { // wg *sync.WaitGroup
			//defer wg.Done()
			err := service.Repository.Store(super)
			log.Println(err)
			if err == nil {
				createdCharacters = append(createdCharacters, super)
				//return nil, err
			}
		}() //&waitGroup

	}

	// waitGroup.Wait()

	return createdCharacters, nil
}

func parseRelatives(relativesResult string) int {
	// TODO handle edge cases
	relatives := strings.Split(relativesResult, "),")
	return len(relatives)
}
