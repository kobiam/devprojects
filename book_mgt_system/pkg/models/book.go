package models

import(
	"github.com/jinzhu/gorm"
	"github.com/devprojects/book_mgt_system"
)

var db *gorm.db

type Book struct{
	gorm.models
	Name string 'gorm:""json:"name"'
	Author string 'json:"author"'
	Publication string 'json:"publication"'
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

