package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../gorm.db")
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
