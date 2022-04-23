package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {

	//dsn := os.Getenv("DB_DSN")
	dsn := "yamanoi:password@tcp(localhost:3306)/go_sample_api"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
