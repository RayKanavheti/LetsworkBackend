package models

import (
	"errors"

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
	ProjectID           uint
	BidderID            uint
}

// CreateBid method creates a bid for a project
func CreateBid(bid Bid) (Bid, error) {

	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		err := db.Create(&bid).Error
		if err == nil {
			return bid, nil
		}
		return bid, errors.New("Unable to create user for session " + err.Error())
	}
	return bid, errors.New("Unable to getdatabase connection")
}
