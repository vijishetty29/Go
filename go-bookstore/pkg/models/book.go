package models

import (
	"go-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
	"github.com/vijishetty29/Go/src/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
