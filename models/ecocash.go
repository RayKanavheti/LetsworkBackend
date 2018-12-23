package models

import (
  "github.com/jinzhu/gorm"
  )
// Ecocash model
type Ecocash struct {
  gorm.Model
  MobileNumber    string
  MerchantCode    string
  FinancialID     uint
}
