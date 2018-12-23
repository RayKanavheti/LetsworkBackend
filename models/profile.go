package models

import (
	"github.com/jinzhu/gorm"
)

// Profile model
type Profile struct {
	gorm.Model
	Name            string `sql:"type:VARCHAR(100)"`
	Surname         string `sql:"type:VARCHAR(50)"`
	PhoneNumber     string `sql:"type:VARCHAR(50)"`
	ImagePath       string `sql:"type:VARCHAR(200)"`
	FacebookProfile string `sql:"type:VARCHAR(150)"`
	TwitterProfile  string `sql:"type:VARCHAR(150)"`
	LinkedInProfile string `sql:"type:VARCHAR(150)"`
}
