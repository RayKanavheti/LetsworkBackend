package models

import (
"github.com/jinzhu/gorm"
)
// Address Model
type Address struct {
  gorm.Model
  Line1   string `sql:"type:VARCHAR(100)"`
  Suburb  string `sql:"type:VARCHAR(30)"`
  City    string `sql:"type:VARCHAR(30)"`
  Country string `sql:"type:VARCHAR(30)"`
 UserID  uint

}
