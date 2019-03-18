package models

import (
	// "database/sql/driver"
	// "encoding/json"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// JSONB type definition
// type JSONB map[string]interface{}

// Project model
type Project struct {
	gorm.Model
	Title        string
	Description  string
	Duration     string
	ProjectFiles []*ProjectFile `gorm:"foreignkey:ProjectID"`
	Status       string         // in-progress, open, completed
	BidEndDate   *time.Time
	Assisted     bool
	Budget       postgres.Jsonb `sql:"type:jsonb"`
	ProjectType  string         // fixed or hourly
	Jobs         []Skill        `gorm:"many2many:project_skills;"`
	Tasks        []*Task        `gorm:"foreignkey:ProjectID"`
	Bids         []*Bid         `gorm:"association_save_reference:false;foreignkey:ProjectID"`
	OwnerID      uint
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

// GetProjectsByOpenProject method get projects by owner id and the status of the project
func GetProjectsByOpenProject(OwnerID int, Status string) ([]Project, error) {
	projects := []Project{}
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		db.Preload("Jobs").Preload("Tasks").Preload("Bids").Preload("ProjectFiles").Where("owner_id = ? AND status = ?", OwnerID, Status).Find(&projects)
		if err == nil {
			return projects, nil
		}
		return []Project{}, errors.New("Unable to get project for session")
	}
	return projects, errors.New("Unable to getdatabase connection")
}

// GetProjectByAssignerID method gets an array of projects based on the User/ AssignerID/ Project OwnerID who cretated it
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
