package repos

import (
	"github.com/jinzhu/gorm"
	"github.com/renato-macedo/superheroapi/domain"
	"log"
)

// Character interface
type CharacterRepository interface {
	Store(user *domain.Character) error
	FindAll() []*domain.Character
	FindByName(name string) ([]*domain.Character, error)
	FindByID(id string) (*domain.Character, error)
	FindByFilter(filter, value string) []*domain.Character
	HasCharacterWhere(filter, value string) bool
}

type CharacterRepositoryDB struct {
	DB *gorm.DB
}

// NewCharacterRepo creates a new instance of the Character repo with the given database connection
func NewCharacterRepo(db *gorm.DB) *CharacterRepositoryDB {
	return &CharacterRepositoryDB{
		DB: db,
	}
}

func (repo *CharacterRepositoryDB) Store(character *domain.Character) error {
	err := repo.DB.Create(character).Error
	return err
}

func (repo *CharacterRepositoryDB) FindByName(name string) ([]*domain.Character, error) {
	var characters []*domain.Character
	err := repo.DB.Where("name LIKE ?", "%"+name+"%").Find(&characters).Error
	if err != nil {
		log.Printf("error querying by name %v\n", err)
		return nil, err
	}
	return characters, nil
}

func (repo *CharacterRepositoryDB) FindByFilter(filter, value string) []*domain.Character {
	var characters []*domain.Character
	repo.DB.Where(filter, value).Find(&characters)
	return characters
}

func (repo *CharacterRepositoryDB) HasCharacterWhere(filter, value string) bool {
	character := &domain.Character{}
	return !repo.DB.Where(filter, value).First(character).RecordNotFound()
}

func (repo *CharacterRepositoryDB) FindAll() []*domain.Character {
	var characters []*domain.Character
	repo.DB.Find(&characters)
	return characters
}

func (repo *CharacterRepositoryDB) FindByID(id string) (*domain.Character, error) {
	character := &domain.Character{}
	err := repo.DB.Where("id = ?", id).First(&character).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return character, nil
}