package repos

import (
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/renato-macedo/superheroapi/domain"
)

// CharacterRepository interface
type CharacterRepository interface {
	Store(user *domain.Character) error
	FindAll() []*domain.Character
	FindByName(name string) ([]*domain.Character, error)
	FindByID(id string) (*domain.Character, error)
	FindByFilter(filter, value string) []*domain.Character
	HasCharacterWhere(filter, value string) bool
	Delete(id string) error
}

// CharacterRepositoryDB implements the CharacterRepository interface for databases
type CharacterRepositoryDB struct {
	DB *gorm.DB
}

// NewCharacterRepo creates a new instance of the Character repo with the given database connection
func NewCharacterRepo(db *gorm.DB) *CharacterRepositoryDB {
	return &CharacterRepositoryDB{
		DB: db,
	}
}

// Store a character in the database
func (repo *CharacterRepositoryDB) Store(character *domain.Character) error {
	err := repo.DB.Create(character).Error
	return err
}

// FindByName exec a LIKE querie with the given name
func (repo *CharacterRepositoryDB) FindByName(name string) ([]*domain.Character, error) {
	var characters []*domain.Character
	err := repo.DB.Where("name_lower_case LIKE ?", "%"+strings.ToLower(name)+"%").Find(&characters).Error
	if err != nil {
		log.Printf("error querying by name %v\n", err)
		return nil, err
	}
	return characters, nil
}

// FindByFilter queries character by the given filter e.g. allignment = ?
func (repo *CharacterRepositoryDB) FindByFilter(filter, value string) []*domain.Character {
	var characters []*domain.Character
	repo.DB.Where(filter, value).Find(&characters)
	return characters
}

// HasCharacterWhere tells if the databaase has a character that satisfies the condition
func (repo *CharacterRepositoryDB) HasCharacterWhere(filter, value string) bool {
	character := &domain.Character{}
	return !repo.DB.Where(filter, value).First(character).RecordNotFound()
}

// FindAll return all records
func (repo *CharacterRepositoryDB) FindAll() []*domain.Character {
	var characters []*domain.Character
	repo.DB.Find(&characters)
	return characters
}

//FindByID return the Character with the given id
func (repo *CharacterRepositoryDB) FindByID(id string) (*domain.Character, error) {
	character := &domain.Character{}
	err := repo.DB.Where("id = ?", id).First(&character).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return character, nil
}

// Delete removes the Character with given id from the database
func (repo *CharacterRepositoryDB) Delete(id string) error {
	character := &domain.Character{
		ID: id,
	}
	err := repo.DB.Delete(character).Error
	return err
}
