package models

import (
"github.com/jinzhu/gorm"
)
// Education model
type Education struct {
  gorm.Model
  SchoolName    string `sql:"type:VARCHAR(50)"`
  EducationType string `sql:"type:VARCHAR(50)"`
  Period        string
  UserID        uint
}
