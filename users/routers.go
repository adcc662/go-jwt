package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-jwt/common"
	"net/http"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", Registration)
	router.POST("/login", UsersLogin)
}

func UserRegister(router *gin.RouterGroup) {
	router.GET("", UserRetrieve)
	router.PUT("", UserUpdate)
}

func Registration(c *gin.Context) {
	userValidator := NewUserValidator()
	if err := userValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&userValidator.user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("my_user", userValidator.user)
	serializer := UserSerializer{C: c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UsersLogin(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	user, err := FindOneUser(&User{Email: loginValidator.userModel.Email})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("not Registered email or invalid password")))
		return
	}

	if user.CheckPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("not Registered email or invalid password")))
		return
	}
	UpdateContextUserModel(c, user.ID)
	serializer := UserSerializer{C: c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func UserRetrieve(c *gin.Context) {
	serializer := UserSerializer{C: c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func UserUpdate(c *gin.Context) {
	myUser := c.MustGet("my_user").(User)
	userValidator := NewUserValidatorFillWith(myUser)
	if err := userValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	userValidator.user.ID = myUser.ID
	if err := myUser.Update(userValidator.user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	UpdateContextUserModel(c, myUser.ID)
	serializer := UserSerializer{C: c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
