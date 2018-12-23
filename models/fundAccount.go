package models

import (
	"github.com/jinzhu/gorm"
)

// FundAccount model
type FundAccount struct {
	gorm.Model
	Amount          float32
	TransactionType string // Either a Deposit or a transfer or maybe a ReleasedFund
	UserID          uint
}
