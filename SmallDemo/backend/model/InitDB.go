package model

import (
	"backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	appConfig := config.GetAppConfig()
	dsn := appConfig.Database.Username + ":" +
		appConfig.Database.Password + "@tcp(" +
		appConfig.Database.IP + ")/" +
		appConfig.Database.Db2 + "?charset=utf8&parseTime=True&loc=Local"
	newDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = newDB
	if err != nil {
		panic("failed to connect database")
	}
}

func getDB() *gorm.DB {
	return db
}
