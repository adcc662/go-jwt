package users

import (
	"errors"
	_ "errors"
	_ "github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"regexp"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"type:varchar(100);unique_index"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(100)"`
	Phone    string `gorm:"type:varchar(100)"`
}

func AutoMigrate() {

}

func (u *User) setPassword(password string) error {
	if len(password) < 6 || len(password) > 12 {
		return errors.New("password must be between 6 and 12 characters long")
	}

	var passwordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[@$&])[a-zA-Z0-9@$&]{6,12}$`)
	if !passwordRegex.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter, " +
			"one lowercase letter, one number and one special character")
	}

	bytePassword := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	return nil

}

func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
