package models

import (
	"crypto/tls"
	"errors"
	"fmt"
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/lithammer/shortuuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

// User model
type User struct {
	gorm.Model
	Username   string `sql:"not null;unique" valid:"Required"`
	Email      string `sql:"type:VARCHAR(50);not null;unique"`
	Password   string
	Profile    Profile      `json:"-" gorm:"save_associations:false"`
	Educations []*Education `json:"-" gorm:"save_associations:false"`
	Address    Address      `json:"-" gorm:"save_associations:false"`
	Financial  Financial    `json:"-" gorm:"save_associations:false"`
	Portfolio  []*Portfolio `json:"-" gorm:"save_associations:false"`
	Skills     []*Skill     `json:"-" gorm:"save_associations:false"`
	UUID       string       `json:"-"`
}

// Validate Email
func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,8}$`)
	return Re.MatchString(email)
}

// SendConfirmationLink
func SendConfirmationLink(userEmail string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "info@letswork.co.zw")
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Confirmation Link")
	m.SetBody("text/html", `<!DOCTYPE html><html><head></head><body><h2>This is a test</h2></body></html>`)

	d := gomail.NewDialer("mail.health263.systems", 25, "ray@health263.systems", "&vXEmW,J3pW]")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("Mail send successfully to %s with password", userEmail)

}

//CreateUser method
func CreateUser(user User) (User, error) {
	if !validateEmail(user.Email) {
		return user, errors.New("Invalid Email")
	}

	user.UUID = shortuuid.New()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPassword)
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		err := db.Save(&user).Error
		if err == nil {
			go SendConfirmationLink(user.Email)
			return user, nil
		}
		return user, errors.New("Unable to create user for session " + err.Error())
	}
	return user, errors.New("Unable to getdatabase connection")
}
