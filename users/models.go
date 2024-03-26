package users

import (
	"errors"
	_ "errors"
	_ "github.com/jinzhu/gorm"
	"go-jwt/common"
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
	db := common.GetDB()
	db.AutoMigrate(&User{})
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

func FindOneUser(condition interface{}) (User, error) {
	db := common.GetDB()
	var model User
	err := db.Where(condition).First(&model).Error
	return model, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func UpdateOne(data interface{}) error {
	db := common.GetDB()
	err := db.Model(data).Update(data).Error
	return err
}

func DeleteOne(data interface{}) error {
	db := common.GetDB()
	err := db.Delete(data).Error
	return err
}

func (u *User) ValidatePhone() error {
	match, _ := regexp.MatchString(`^\d{10}$`, u.Phone)
	if !match {
		return errors.New("phone must be exactly 10 digits")
	}
	return nil
}

func (u *User) ValidateEmail() error {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, u.Email)
	if !match {
		return errors.New("email is invalid")
	}
	return nil
}
