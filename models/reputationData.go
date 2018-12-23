package models

import (
  "github.com/jinzhu/gorm"
)
// ReputationData model
type ReputationData struct {
  gorm.Model
  OnTime         float32 `sql:"DEFAULT:0.00"`
  OnBudget       float32  `sql:"DEFAULT:0.00"`
  CategoryRating CategoryRating
  ReviewID       uint

}
