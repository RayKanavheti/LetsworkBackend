package models

import (
  "github.com/jinzhu/gorm"
)

// Bank Model
type Bank struct {
  gorm.Model
  BankName       string `sql:"type:VARCHAR(100)"`
  AccountName    string `sql:"type:VARCHAR(50)"`
  AccountNumber  string `sql:"type:VARCHAR(25)"`
  BranchCode     string `sql:"type:VARCHAR(25)"`
  BranchName     string `sql:"type:VARCHAR(50)"`
  FinancialID     uint
}
