package character

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/renato-macedo/superheroapi/database"
	"github.com/renato-macedo/superheroapi/superhero"
)

func init() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

type TestPair struct {
	input  string
	output int
}

func TestService_Create(t *testing.T) {
	db := database.Connect("test")
	repository := NewCharacterRepo(db)

	characterService := &Service{
		Repository: repository,
		API: &superhero.Service{
			BaseURL: "https://superheroapi.com/api",
			APIKey:  os.Getenv("API_KEY"),
		},
	}

	TestCases := []TestPair{
		{
			input:  "Spider-Man",
			output: 3,
		},
		{
			input:  "Spider-Man",
			output: 0,
		},
		{
			input:  "Non Existing Super",
			output: -1, // should be error
		},
	}

	for _, testCase := range TestCases {
		result, err := characterService.Create(testCase.input)

		if testCase.output == -1 {
			if err == nil {
				t.Errorf("Test case expected error but got |%v|", err)

			}
			continue
		}

		createdCharactersCount := len(result)

		if createdCharactersCount != testCase.output {
			t.Errorf("Expected result for input |%v| to be |%v| but instead got %v", testCase.input, testCase.output, createdCharactersCount)
		}
	}

	t.Cleanup(func() {
		repository.DB.DropTable(&Character{})
	})

}
