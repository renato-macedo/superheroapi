package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/renato-macedo/superheroapi/character"
	"github.com/renato-macedo/superheroapi/database"
	"github.com/renato-macedo/superheroapi/superhero"
	"github.com/stretchr/testify/assert"
)

func TestGETSupers(t *testing.T) {

	cleanup := setup(true)

	TestsCases := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError        bool
		expectedCode         int
		expectedResultsCount int
	}{
		{
			description:          "GET /super",
			route:                "/super",
			expectedError:        false,
			expectedCode:         200,
			expectedResultsCount: 6,
		},
		{
			description:          "GET /super/heros",
			route:                "/super/heros",
			expectedError:        false,
			expectedCode:         200,
			expectedResultsCount: 1,
		},
		{
			description:          "GET /super/villains",
			route:                "/super/villains",
			expectedError:        false,
			expectedCode:         200,
			expectedResultsCount: 5,
		},
		{
			description:          "GET /super/search?name=captain",
			route:                "/super/search?name=captain",
			expectedError:        false,
			expectedCode:         200,
			expectedResultsCount: 0,
		},
		{
			description:          "GET /super/search",
			route:                "/super/search",
			expectedError:        false,
			expectedCode:         400,
			expectedResultsCount: 0,
		},
	}

	app := startApp()

	// Iterate through test single test cases
	for _, test := range TestsCases {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		if test.expectedCode < 400 {
			// parse sucessul response
			body, err := parseGETResponse(res.Body)

			assert.Nilf(t, err, test.description)

			assert.Equalf(t, test.expectedResultsCount, len(body), test.description)
		} else {
			// parse error response, maybe it would be better to add a Test function just for bad requests
			body, err := parseErrorResponse(res.Body)

			assert.Nilf(t, err, test.description)
			assert.NotEmptyf(t, body["message"], test.description)
		}

	}

	cleanup()

}

// -------------------------------------------------------------

// ----------------------------------------------------------------

func TestPOSTSupers(t *testing.T) {
	cleanup := setup(false)
	TestsCases := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError          bool
		expectedCode           int
		payload                string
		expectedCharacterCount int
	}{
		{
			description:            "POST /super -d \"{\"name\": \"mister\"}\" ",
			route:                  "/super",
			expectedError:          false,
			expectedCode:           201,
			payload:                `{"name": "mister"}`,
			expectedCharacterCount: 6,
		},
		{
			description:            "POST /super -d \"{\"hero\": \"mister\"}\" ",
			route:                  "/super",
			expectedError:          false,
			expectedCode:           400,
			payload:                `{"hero": "man"}`,
			expectedCharacterCount: 0,
		},
	}

	app := startApp()

	for _, test := range TestsCases {

		req, _ := http.NewRequest(
			"POST",
			test.route,
			bytes.NewBuffer([]byte(test.payload)),
		)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len([]byte(test.payload))))

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := parsePOSTResponse(res.Body)

		assert.Nilf(t, err, test.description)

		assert.Equalf(t, test.expectedCharacterCount, body.CharactersCount, test.description)

	}

	cleanup()

}

func setup(populate bool) func() {
	os.Setenv("ENV", "test")
	db := database.Connect("test")
	repository := character.NewCharacterRepo(db)

	if populate {
		service := character.Service{
			Repository: repository,
			API: &superhero.Service{
				BaseURL: "https://superheroapi.com/api",
				APIKey:  os.Getenv("API_KEY"),
			},
		}

		_, err := service.Create("Mister")
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	// cleanup
	return func() {
		repository.DB.Delete(&character.Character{})
	}
}

func parseGETResponse(body io.Reader) ([]character.DTO, error) {
	var response []character.DTO
	err := json.NewDecoder(body).Decode(&response)

	return response, err
}

func parsePOSTResponse(body io.Reader) (character.CreatedResponse, error) {

	response := new(character.CreatedResponse)
	resp, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(resp)
	}
	err = json.Unmarshal(resp, response)
	if err != nil {
		log.Println(resp)
	}
	return *response, err
}

func parseErrorResponse(body io.Reader) (map[string]string, error) {
	// if error try to parse with map of strings
	generic := make(map[string]string)
	err := json.NewDecoder(body).Decode(&generic)
	return generic, err
}
