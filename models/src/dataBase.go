package src

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//connectDB
func connectDB() *gorm.DB {
	dsn := "root:123456789@tcp(localhost:3306)/go_training?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connect successfull")
	}
	return db
}
