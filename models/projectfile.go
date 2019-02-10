package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// ProjectFile struct
type ProjectFile struct {
	gorm.Model
	Path      string `sql:"type:VARCHAR(250)"`
	ProjectID int
}

// CreateFiles method creates multiple Files for the user
func CreateFiles(projectFiles []ProjectFile) ([]ProjectFile, error) {
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		for _, projectFile := range projectFiles {
			err := db.Save(&projectFile).Error
			if err != nil {
				return projectFiles, errors.New("Unable to create File for session" + err.Error())
			}
		}
		if err == nil {
			return projectFiles, nil
		}
		return projectFiles, errors.New("Unable to create File" + err.Error())
	}
	return projectFiles, errors.New("Unable to get database connection")
}
