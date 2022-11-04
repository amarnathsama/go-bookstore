package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "input dsn here"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
