package main

import (
	_ "fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"go-jwt/users"
	_ "go-jwt/users"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

func main() {

}
