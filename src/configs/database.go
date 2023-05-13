package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateDatabaseConnection() {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		GetEnv("DB_USER"),
		GetEnv("DB_PASS"),
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"),
		GetEnv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connString), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Error connecting to database!")
	}

	DB.AutoMigrate()
}
