package models

import (
	"github.com/jinzhu/gorm"
)
// Skill model
type Skill struct {
	gorm.Model
  Title     string
}
