package character

// Character is a Hero or a Villain
type Character struct {
	ID                string `json:"id" gorm:"type:uuid;primary_key"`
	APIID             string `json:"-" gorm:"type:int;unique_index"`
	Name              string `json:"name" gorm:"type:varchar(100)"`
	NameLowerCase     string `json:"-" gorm:"type:varchar(100)"`
	FullName          string `json:"full-name" gorm:"type:varchar(100)"`
	Intelligence      string `json:"intelligence" gorm:"type:varchar(50)"`
	Power             string `json:"power" gorm:"type:varchar(50)"`
	Occupation        string `json:"occupation,omitempty" gorm:"type:varchar(400)"`
	Image             string `json:"image,omitempty" gorm:"type:varchar(250)"`
	Alignment         string `json:"alignment" gorm:"type:varchar(50)"`
	GroupAffiliation  string `json:"-" gorm:"type:varchar(400)"`
	NumberOfRelatives int    `json:"number_of_relatives"`
}
