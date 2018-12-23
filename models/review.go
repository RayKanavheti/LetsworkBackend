package models

import (
  "github.com/jinzhu/gorm"
)
// Review model
type Review struct {
  gorm.Model
  Comment           string
  Role              string // Either Employer or Employee
  FromUserID        uint
  ToUserID          uint
  ReputationData    ReputationData
}
