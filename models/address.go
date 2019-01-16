package models

import (
"github.com/jinzhu/gorm"
"errors"
)
// Address Model
type Address struct {
  gorm.Model
  Line1   string `sql:"type:VARCHAR(100)"`
  Suburb  string `sql:"type:VARCHAR(30)"`
  City    string `sql:"type:VARCHAR(30)"`
  Country string `sql:"type:VARCHAR(30)"`
  UserID  uint

}

// CreateAddress method creates address for the user
func CreateAddress(userAddress Address) (Address, error) {

	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		err := db.Save(&userAddress).Error
		if err == nil {
			return userAddress, nil
		}
		return userAddress, errors.New("Unable to create address" + err.Error())
	}
	return userAddress, errors.New("Unable to getdatabase connection")
}
