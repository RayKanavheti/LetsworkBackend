package models

import (
	"github.com/jinzhu/gorm"
)

// Financial model
type Financial struct {
	gorm.Model
	Bank    Bank
	Ecocash Ecocash
	UserID  uint

}
