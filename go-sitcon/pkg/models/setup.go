package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Admin@123@tcp(localhost:3306)/prepim-go?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println("db", db)
	fmt.Println("error", err)

	fmt.Println("DB Connection Established", db)

	DB = db
}
