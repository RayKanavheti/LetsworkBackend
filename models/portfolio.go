package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Portfolio model
type Portfolio struct {
	gorm.Model
	Name   string `sql:"type:VARCHAR(100)"`
	Link   string `sql:"type:VARCHAR(50)"`
	UserID uint
}

// CreatePortfolio method creates portfolio for the user
func CreatePortfolio(userPortfolio Portfolio) (Portfolio, error) {

	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		err := db.Save(&userPortfolio).Error
		if err == nil {
			return userPortfolio, nil
		}
		return userPortfolio, errors.New("Unable to create portfolio" + err.Error())
	}
	return userPortfolio, errors.New("Unable to get database connection")
}
