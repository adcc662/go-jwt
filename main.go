package main

import (
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-jwt/common"
	"go-jwt/users"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

// @title Swagger Go JWT API
// @version 1.0
// @description This is a sample server for Go JWT.
// @host localhost:8080
// @BasePath /api
func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	fmt.Print("Server started at localhost:8080")
	r.Run()

}
