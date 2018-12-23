package models

import (
	"github.com/jinzhu/gorm"
)

// Portfolio model
type Portfolio struct {
	gorm.Model
	Name   string `sql:"type:VARCHAR(100)"`
	Link   string `sql:"type:VARCHAR(50)"`
	UserID uint
}
