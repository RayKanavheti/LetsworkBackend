package models

import (
	"github.com/jinzhu/gorm"
)

// CategoryRating model
type CategoryRating struct {
	gorm.Model
	Communication        float32 `sql:"DEFAULT:0.00"`
	Expertise            float32 `sql:"DEFAULT:0.00"`
	Quality              float32 `sql:"DEFAULT:0.00"`
	HireAgain            float32 `sql:"DEFAULT:0.00"`
	Professionalism      float32 `sql:"DEFAULT:0.00"`
	ClaritySpecification float32 `sql:"DEFAULT:0.00"`
	PaymentPromise       float32 `sql:"DEFAULT:0.00"`
	WorkForAgain         float32 `sql:"DEFAULT:0.00"`
	ReputationDataID     uint
}
