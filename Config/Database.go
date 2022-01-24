package Config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func Init() {

	dbUser := EnvConfigs["DB_USER"]
	dBPassword := EnvConfigs["DB_PASSWORD"]
	dbHost := EnvConfigs["DB_HOST"]
	dbPort := EnvConfigs["DB_PORT"]
	dbName := EnvConfigs["DB_NAME"]

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dBPassword, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
}

func CloseDB() {
	return
}
