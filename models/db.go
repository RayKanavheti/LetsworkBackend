package models

import (
  "log"
  "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/postgres" // going to be used
   "os"
   "errors"
  )

  // Postgres
func getDBConnection() (*gorm.DB, error) {
db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=letswork password=#LetMeIn sslmode=disable")
	db.DB().SetMaxIdleConns(10)
	if os.Getenv("GO_ENV") != "tests" {
		db.LogMode(true)
	}
	return db, err
}

//InitDB for initilaizing DB operations
func InitDB() {
	err := autoMigrateTables()
	if err != nil {
		log.Fatal(err)
		return
	}
}

//AutoMigrateTables to be run at startup. Its a bit more resource hungry and slows startup.
func autoMigrateTables() error {
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
    db.SingularTable(true)
		db.Debug().AutoMigrate(&User{}, &Profile{}, &Address{}, &Education{}, &Portfolio{}, &Financial{}, &Bank{}, &Ecocash{},
			&WorkDone{}, &Project{}, &ProjectCategory{}, &Bid{}, &Review{}, &ReputationData{}, &Task{}, &CategoryRating{}, &OverallRating{}, &Skill{},
			&Notification{}, &ReleasedFund{}, &FundAccount{})
		if err == nil {
			return nil
		}
		return errors.New("Unable to Automigrate tables " + err.Error())
	}
	// return errors.New("Unable to getdatabase connection")
  return errors.New(err.Error())
}
