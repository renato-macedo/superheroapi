package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CharacterResponse represents the Character json object returned from the SuperHeroAPI
type CharacterResponse struct {
	// Response    string `json:"response"`
	ID string `json:"id"`

	Name string `json:"name"`

	Powerstats struct {
		Intelligence string `json:"intelligence"`
		Strength     string `json:"strength"`
		Speed        string `json:"speed"`
		Durability   string `json:"durability"`
		Power        string `json:"power"`
		Combat       string `json:"combat"`
	} `json:"powerstats" `

	Biography struct {
		FullName        string   `json:"full-name"`
		AlterEgos       string   `json:"alter-egos"`
		Aliases         []string `json:"aliases"`
		PlaceOfBirth    string   `json:"place-of-birth"`
		FirstAppearance string   `json:"first-appearance"`
		Publisher       string   `json:"publisher"`
		Alignment       string   `json:"alignment"`
	} `json:"biography"`

	Appearance struct {
		Gender    string   `json:"gender"`
		Race      string   `json:"race"`
		Height    []string `json:"height"`
		Weight    []string `json:"weight"`
		EyeColor  string   `json:"eye-color"`
		HairColor string   `json:"hair-color"`
	} `json:"appearance"`

	Work struct {
		Occupation string `json:"occupation"`
		Base       string `json:"base"`
	} `json:"work"`

	Connections struct {
		GroupAffiliation string `json:"group-affiliation"`
		Relatives        string `json:"relatives"`
	} `json:"connections"`

	Image struct {
		URL string `json:"url"`
	} `json:"image"`
}

// SearchResult represents the Character json object returned from the SuperHeroAPI
type SearchResult struct {
	Response   string               `json:"response"`
	ResultsFor string               `json:"results-for"`
	Results    []*CharacterResponse `json:"results"`
}

// SuperHeroAPIService need a baseURL and a API key
type SuperHeroAPIService struct {
	BaseURL string
	APIKey  string
}

// SearchCharacter call search endpoint
func (api *SuperHeroAPIService) SearchCharacter(name string) (*SearchResult, error) {

	resp, err := http.Get(fmt.Sprintf("%v/%v/search/%v", api.BaseURL, api.APIKey, name))
	if err != nil {
		log.Printf("Error search for character %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body %v\n", err)
		return nil, err
	}
	//log.Println(string(body))
	response := &SearchResult{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Error unmarshalling json %v\n", err)
		return nil, err
	}
	return response, nil
}
