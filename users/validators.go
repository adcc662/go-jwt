package users

import (
	"github.com/gin-gonic/gin"
	"go-jwt/common"
)

type UserValidator struct {
	UserV struct {
		Username string `form:"username" json:"username" binding:"required"`
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
		Phone    string `form:"phone" json:"phone" binding:"required"`
	}
	user User `json:"-"`
}

func (self *UserValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.user.Username = self.UserV.Username
	self.user.Email = self.UserV.Email
	self.user.Phone = self.UserV.Phone

	if self.UserV.Password != "" {
		if err := self.user.setPassword(self.UserV.Password); err != nil {
			return err
		}
	}

	if self.UserV.Phone != "" {
		if err := self.user.ValidatePhone(); err != nil {
			return err
		}
	}

	if self.UserV.Email != "" {
		if err := self.user.ValidateEmail(); err != nil {
			return err
		}
	}

	return nil
}

func NewUserValidator() UserValidator {
	userValidator := UserValidator{}
	//userModelValidator.User.Email ="w@g.cn"
	return userValidator
}

func NewUserValidatorFillWith(user User) UserValidator {
	userValidator := NewUserValidator()
	userValidator.UserV.Username = user.Username
	userValidator.UserV.Email = user.Email
	userValidator.UserV.Password = common.NBRandomPassword

	return userValidator
}

type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password"json:"password" binding:"exists,min=6,max=12"`
	} `json:"user"`
	userModel User `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.userModel.Email = self.User.Email
	return nil
}

// NewLoginValidator You can put the default value of a Validator here
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
