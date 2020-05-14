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
	FindByName(name string) (*domain.Character, error)
	FindByID(id string) (*domain.Character, error)
	FindByFilter(filter, value string) []*domain.Character
	FindFirstByFilter(filter, value string) *domain.Character
}

type CharacterRepositoryDB struct {
	db *gorm.DB
}

// NewCharacterRepo creates a new instance of the Character repo with the given database connection
func NewCharacterRepo(db *gorm.DB) *CharacterRepositoryDB {
	return &CharacterRepositoryDB{
		db: db,
	}
}

func (repo *CharacterRepositoryDB) Store(character *domain.Character) error {
	err := repo.db.Create(character).Error
	return err
}

func (repo *CharacterRepositoryDB) FindByName(name string) []*domain.Character {
	var characters []*domain.Character
	err := repo.db.Where("name LIKE ?", "%"+name+"%").Find(&characters).Error
	if err != nil {
		log.Printf("error querying by name %v\n", err)
	}
	return characters
}

func (repo *CharacterRepositoryDB) FindByFilter(filter, value string) []*domain.Character {
	var characters []*domain.Character
	repo.db.Where(filter, value).Find(&characters)
	return characters
}
func (repo *CharacterRepositoryDB) FindFirstByFilter(filter, value string) *domain.Character {
	character := &domain.Character{}
	repo.db.Where(filter, value).First(character)
	return character
}
func (repo *CharacterRepositoryDB) HasCharacterWhere(filter, value string) bool {
	character := &domain.Character{}
	return !repo.db.Where(filter, value).First(character).RecordNotFound()
}
