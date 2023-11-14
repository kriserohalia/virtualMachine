package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/virtual_machine"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{},&Employee{})

	DB = db
}
