package models

import (
"github.com/jinzhu/gorm"
"errors"
)
// Education model
type Education struct {
  gorm.Model
  SchoolName    string `sql:"type:VARCHAR(50)"`
  EducationType string `sql:"type:VARCHAR(50)"`
  Period        string
  UserID        uint
}
// CreateEducation method creates address for the user
func CreateEducation(userEducation Education) (Education, error) {

	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		err := db.Save(&userEducation).Error
		if err == nil {
			return userEducation, nil
		}
		return userEducation, errors.New("Unable to create education" + err.Error())
	}
	return userEducation, errors.New("Unable to get database connection")
}
