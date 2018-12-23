package models

import (
	"github.com/jinzhu/gorm"
)

// OverallRating model
type OverallRating struct {
	gorm.Model
	AvaregeRating float32
	CountRaters   int
	TotalRatings  float32
	UserID        uint
}
