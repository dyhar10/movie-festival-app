package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/dyhar10/movie-festival-app/configs"
)

func ConfigDB() *gorm.DB {
	USER := configs.GetEnv("DB_USER")
	PASS := configs.GetEnv("DB_PASSWORD")
	HOST := configs.GetEnv("DB_HOST")
	PORT := configs.GetEnv("DB_PORT")
	DBNAME := configs.GetEnv("DB_NAME")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
