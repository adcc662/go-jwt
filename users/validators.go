package users

import (
	"github.com/gin-gonic/gin"
	"go-jwt/common"
)

type UserValidator struct {
	UserV struct {
		Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=6,max=12"`
		Phone    string `form:"phone" json:"phone" binding:"exists,numeric"`
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
		self.user.setPassword(self.UserV.Password)
	}

	if self.UserV.Phone != "" {
		self.user.ValidatePhone()
	}

	if self.UserV.Email != "" {
		self.user.ValidateEmail()
	}

	return nil
}
