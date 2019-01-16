package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Project model
type Project struct {
	gorm.Model
	Title      string
	Desciption string
	DocPath1   string
	DocPath2   string
	DocPath3   string
	Duration   string
	IsComplete bool
	Status     string
	Assisted   string
	Budget     struct {
		minimum float32
		maximum float32
	} `sql:"type:jsonb"`
	ProjectType string // fixed or hourly
	Project     ProjectCategory
	Tasks       []*Task
	OwnerID     uint
}

// CreateProject method creates a project
func CreateProject(userProject Project) (Project, error) {

	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		err := db.Save(&userProject).Error
		if err == nil {
			return userProject, nil
		}
		return userProject, errors.New("Unable to create user for session " + err.Error())
	}
	return userProject, errors.New("Unable to getdatabase connection")
}

// UpdateProject method updates a project
func UpdateProject(userProject Project) (Project, error) {

	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		err := db.Save(&userProject).Error
		if err == nil {
			return userProject, nil
		}
		return userProject, errors.New("Unable to create user for session " + err.Error())
	}
	return userProject, errors.New("Unable to getdatabase connection")
}
