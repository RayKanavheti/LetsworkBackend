package models

import (
	"github.com/jinzhu/gorm"
)
// ReleasedFund model
type ReleasedFund struct {
	gorm.Model
	Amount       float32
	FromUserID   uint
  ToUserID     uint
}
