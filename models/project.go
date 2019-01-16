package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/jinzhu/gorm"
)

// JSONB type definition
type JSONB map[string]interface{}

// Project model
type Project struct {
	gorm.Model
	Title       string
	Description  string
	DocPath1    string
	DocPath2    string
	DocPath3    string
	Duration    string
	IsComplete  bool
	Status      string
	Assisted    bool
	Budget      JSONB             `sql:"type:jsonb"`
	ProjectType string            // fixed or hourly
	Jobs        []ProjectCategory `gorm:"many2many:job_projcat;"`
	Tasks       []*Task
	OwnerID     uint
}

// Value method
func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}
// Scan method
func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

// CreateProject method creates a project
func CreateProject(userProject Project) (Project, error) {

	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		err := db.Create(&userProject).Error
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
