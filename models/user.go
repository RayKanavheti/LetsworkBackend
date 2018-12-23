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
	ResetKey   string       `json:"-"`
	Verified   bool
}

// ValidateEmail validates an email received
func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,8}$`)
	return Re.MatchString(email)
}

// SendConfirmationLink method sends a link to the user's email with a uuid generated code for verification
func SendConfirmationLink(userEmail string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "ray@health263.systems")
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Confirmation Link")
	m.SetBody("text/html", `<!DOCTYPE html><html><head></head><body><h2>This is a test</h2></body></html>`)

	d := gomail.NewDialer("mail.health263.systems", 25, "ray@health263.systems", "Raycanas199425%")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("Mail send successfully to %s", userEmail)

}

//CreateUser method creates a new user
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

//GetUsers function. Lists all Users with full details
func GetUsers() ([]User, error) {
	user := []User{}
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		db.Find(&user)
		if err == nil {
			return user, nil
		}
		return []User{}, errors.New("Unable to get user for session")
	}
	return user, errors.New("Unable to getdatabase connection")
}

//GetUserByID getting the details of user using the user ID.
func GetUserByID(id int) (User, error) {
	user := User{}
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		db.Where("id = ?", id).Find(&user)
		if err == nil {
			if user.ID == 0 {
				return user, errors.New("Unable to get user for session")
			}
			return user, nil
		}
		return User{}, errors.New("Unable to get user for session")
	}
	return user, errors.New("Unable to getdatabase connection")
}

//GetUserByUUID getting the details of the user using the uuid inorder to verify the user's email.
func GetUserByUUID(uuid string) (User, error) {
	user := User{}
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		db.Where("uuid = ?", uuid).Find(&user)
		if err == nil {
			if user.ID == 0 {
				return user, errors.New("Unable to get user with that uuid")
			}
			return user, nil
		}
		return User{}, errors.New("Unable to get user for session")
	}
	return user, errors.New("Unable to getdatabase connection")
}

// GetUserByEmail getting the details of the user using their email address
func GetUserByEmail(userEmail string) (User, error) {
	user := User{}
	if !validateEmail(userEmail) {
		return user, errors.New("Invalid Email")
	}
	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		db.Where("email = ?", userEmail).Find(&user)
		if err == nil {
			if user.ID == 0 {
				return user, errors.New("Unable to get user with email")
			}
			return user, nil
		}
		return User{}, errors.New("Unable to get user with this email")
	}
	return user, errors.New("Unable to getdatabase connection")
}

//UpdateVerificationField sets the verified field to true after confirmation link has been clicked
func UpdateVerificationField(user User) (User, error) {

	db, err := getDBConnection()
	defer db.Close()
	if err == nil {
		db.Model(&user).Updates(map[string]interface{}{"verified": true, "uuid": nil})
		if err == nil {
			return user, nil
		}
		return User{}, errors.New("Unable to update user")
	}
	return user, errors.New("Unable to getdatabase connection")
}

// ResetPassword sends a link to reset password
func ResetPassword(userEmail string) (string) {
	 status := SendResetPasswordlink(userEmail)
	 	fmt.Printf("status %s", status)
	if status == "error" {
			return "error"
	}
	return "success"
}
// SendResetPasswordlink sends a link to reset password
func SendResetPasswordlink(userEmail string) (string){
	m := gomail.NewMessage()
	m.SetHeader("From", "ray@health263.systems")
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Request to reset password from Let's work")
	m.SetBody("text/html", `<!DOCTYPE html><html><head></head><body><h2>This is a test for password sending</h2></body></html>`)

	d := gomail.NewDialer("mail.health263.systems", 25, "ray@health263.systems", "Raycanas199425%")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf(err.Error())
		return "error"
	}
	fmt.Printf("Mail send successfully to %s", userEmail)
	return "success"
}
