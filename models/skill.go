package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Skill model
type Skill struct {
	gorm.Model
	Title string
}

// CreateSkills method creates a new user
func CreateSkills(skills []Skill) ([]Skill, error) {
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		for _, skill := range skills {
			err := db.Create(&skill).Error
			if err != nil {
				return skills, errors.New("Unable to create skill for session" + err.Error())
			}
		}
		if err == nil {
			return skills, nil
		}
		return skills, errors.New("Unable to create skill" + err.Error())
	}
	return skills, errors.New("Unable to get database connection")
}
// GetSkills method retrieves all skills
func GetSkills()([]Skill, error) {
	skills := []Skill{}
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		db.Find(&skills)
		if err == nil {
			return skills, nil
		}
		return []Skill{}, errors.New("Unable to get user for session")
	}
	return skills, errors.New("Unable to getdatabase connection")
}
