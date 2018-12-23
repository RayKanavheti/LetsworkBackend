package models

import (
	"github.com/jinzhu/gorm"
)

// Bid model
type Bid struct {
	gorm.Model
	Amount              string `sql:"type:VARCHAR(100)"`
	Retracted           bool
	Description         string  `sql:"type:VARCHAR(300)"`
	MilestonePercentage float32 `sql:"DEFAULT:0.00"`
	Awarded             bool
	DurationInDays      string `sql:"type:VARCHAR(3)"`
}
