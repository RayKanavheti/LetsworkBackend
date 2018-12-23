package models

import (
  "github.com/jinzhu/gorm"
)
// ProjectCategory model
type ProjectCategory struct {
  gorm.Model
  Name  string `sql:"type:VARCHAR(60)"`
}
