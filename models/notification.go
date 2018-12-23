package models

import(
  "github.com/jinzhu/gorm"
)
// Notification model
type Notification struct {
  gorm.Model
  message   string
  read      bool
}
