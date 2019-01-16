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
// CreateEducations method creates multiple educations for the user
func CreateEducations (userEducations []Education) ([]Education, error) {
  db, err := getDBConnection()
  defer db.Close()
  if err == nil {
    for _, userEducation := range userEducations {
      err := db.Save(&userEducation).Error
      if err != nil {
				return userEducations, errors.New("Unable to create education for session" + err.Error())
			}
    }
    if err == nil {
			return userEducations, nil
		}
      return userEducations, errors.New("Unable to create education" + err.Error())
  }
  	return userEducations, errors.New("Unable to get database connection")
}
