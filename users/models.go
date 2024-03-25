package users

import (
	_ "errors"
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"type:varchar(100);unique_index"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(100)"`
	Phone    string `gorm:"type:varchar(100)"`
}
