package models

import (
	"errors"

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
	UserID          uint

}

// CreateProfile method creates a new user
func CreateProfile(userProfile Profile) (Profile, error) {

	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		err := db.Save(&userProfile).Error
		if err == nil {
			return userProfile, nil
		}
		return userProfile, errors.New("Unable to create user for session " + err.Error())
	}
	return userProfile, errors.New("Unable to getdatabase connection")
}
