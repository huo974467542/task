package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"task4/model"
)

var DB *gorm.DB

func Initdb() *gorm.DB {
	dsn := "root:root@tcp(192.168.232.144:3306)/go-test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	err = db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		log.Fatal("failed to migrate db")
	}
	db = db.Debug()
	return db
}
