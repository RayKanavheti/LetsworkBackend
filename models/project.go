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
	Description string
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
	Tasks       []*Task           `gorm:"foreignkey:ProjectID"`
	Bids        []*Bid            `gorm:"association_save_reference:false;foreignkey:ProjectID"`
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

// CreateProject method creates a new project
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

// GetProjectByID method gets the project by its ID
func GetProjectByID(id int) (Project, error) {
	project := Project{}
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		db.Preload("Jobs").Preload("Tasks").Where("id = ?", id).Find(&project)
		if err == nil {
			if project.ID == 0 {
				return project, errors.New("Unable to get project for session")
			}
			return project, nil
		}
		return Project{}, errors.New("Unable to get project for session")
	}
	return project, errors.New("Unable to getdatabase connection")
}

// GetProjectByAssignerID method gets an array of projects based on the User/ AssignerID/ Project OwnerID
func GetProjectByAssignerID(AssignerID int) ([]Project, error) {
	projects := []Project{}
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		db.Preload("Jobs").Preload("Tasks").Where("owner_id = ?", AssignerID).Find(&projects)
		if err == nil {
			return projects, nil
		}
		return []Project{}, errors.New("Unable to get user for session")
	}
	return projects, errors.New("Unable to getdatabase connection")
}
