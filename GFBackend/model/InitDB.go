package model

import (
	"GFBackend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	appConfig := config.AppConfig
	dsn := appConfig.Database.Username + ":" +
		appConfig.Database.Password + "@tcp(" +
		appConfig.Database.IP + ")/" +
		appConfig.Database.DB1 + "?charset=utf8&parseTime=True&loc=Local"
	newDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	DB = newDB
	if err != nil {
		panic("failed to connect database")
	}
}
