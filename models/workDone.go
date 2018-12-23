package models

import (
	"github.com/jinzhu/gorm"
)

// WorkDone model
type WorkDone struct {
	gorm.Model
	FileName   string
	PathToFile string
	LinkToWrk  string
	Filesize   string
	ToUserID   uint
	FromUserID uint
	ProjectID  uint
}
