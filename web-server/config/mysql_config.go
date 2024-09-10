package config

import (
	"fmt"
	"log"
	"web-server/models/member"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
    log.Println("Initializing database connection...")
    Connect()
}

func Connect(){
	dsn := "root:test@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local"
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})



    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

	fmt.Println("connected")

    err = db.AutoMigrate(&member.Member{})
    if err != nil {
        log.Fatalf("Error during AutoMigrate: %v", err)
    }
}

func GetDB()(*gorm.DB){  
	if db == nil {
        log.Fatal("Database connection is not initialized")
    }
    return db
}