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

// CreatePortfolios method creates multiple Portfolios for the user
func CreatePortfolios (userPortfolios []Portfolio) ([]Portfolio, error) {
  db, err := getDBConnection()
  defer db.Close()
  if err == nil {
    for _, userPortfolio := range userPortfolios{
      err := db.Save(&userPortfolio).Error
      if err != nil {
				return userPortfolios, errors.New("Unable to create portfolio for session" + err.Error())
			}
    }
    if err == nil {
			return userPortfolios, nil
		}
      return userPortfolios, errors.New("Unable to create Portfolio" + err.Error())
  }
  	return userPortfolios, errors.New("Unable to get database connection")
}
