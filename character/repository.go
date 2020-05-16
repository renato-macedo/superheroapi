package character

import (
	"errors"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {
	Store(user *Character) error
	FindAll() []Character
	FindByName(name string) ([]Character, error)
	FindByID(id string) (*Character, error)
	FindByFilter(filter, value string) []Character
	HasCharacterWhere(filter, value string) bool
	Delete(id string) error
}

// RepositoryDB implements the Repository interface for databases
type RepositoryDB struct {
	DB *gorm.DB
}

// NewCharacterRepo creates a new instance of the Character repo with the given database connection
func NewCharacterRepo(db *gorm.DB) *RepositoryDB {
	return &RepositoryDB{
		DB: db,
	}
}

// Store a character in the database
func (repo *RepositoryDB) Store(character *Character) error {
	err := repo.DB.Create(character).Error
	return err
}

// FindByName exec a LIKE querie with the given name
func (repo *RepositoryDB) FindByName(name string) ([]Character, error) {
	var characters []Character
	err := repo.DB.Where("name_lower_case LIKE ?", "%"+strings.ToLower(name)+"%").Find(&characters).Error
	if err != nil {
		log.Printf("error querying by name %v\n", err)
		return nil, err
	}
	return characters, nil
}

// FindByFilter queries character by the given filter e.g. allignment = ?
func (repo *RepositoryDB) FindByFilter(filter, value string) []Character {
	var characters []Character
	repo.DB.Where(filter, value).Find(&characters)
	return characters
}

// HasCharacterWhere tells if the databaase has a character that satisfies the condition
func (repo *RepositoryDB) HasCharacterWhere(filter, value string) bool {
	character := &Character{}
	return !repo.DB.Where(filter, value).First(character).RecordNotFound()
}

// FindAll return all records
func (repo *RepositoryDB) FindAll() []Character {
	var characters []Character
	repo.DB.Find(&characters)
	return characters
}

//FindByID return the Character with the given id
func (repo *RepositoryDB) FindByID(id string) (*Character, error) {
	character := &Character{}
	err := repo.DB.Where("id = ?", id).First(&character).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return character, nil
}

// Delete removes the Character with given id from the database
func (repo *RepositoryDB) Delete(id string) error {
	character := &Character{
		ID: id,
	}
	result := repo.DB.Delete(character)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Character does not exist")
	}

	return nil
}
