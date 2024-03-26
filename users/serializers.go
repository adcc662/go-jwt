package users

import (
	"github.com/gin-gonic/gin"
)

type ProfileSerializer struct {
	C *gin.Context
	User
}

type ProfileResponse struct {
	ID       uint   `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (self *ProfileSerializer) Response() ProfileResponse {
	return ProfileResponse{
		ID:       self.ID,
		Username: self.Username,
		Email:    self.Email,
		Phone:    self.Phone,
	}

}

type UserSerializer struct {
	C *gin.Context
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (self *UserSerializer) Response() UserResponse {
	return UserResponse{
		Username: self.C.PostForm("username"),
		Email:    self.C.PostForm("email"),
		Phone:    self.C.PostForm("phone"),
	}
}
