package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() *gorm.DB {
	dbConfig := struct {
		Host   string
		User   string
		Pwd    string
		DbName string
		Port   string
	}{
		Host:   os.Getenv("DB_HOST"),
		User:   os.Getenv("DB_USER"),
		Pwd:    os.Getenv("DB_PWD"),
		DbName: os.Getenv("DB_NAME"),
		Port:   os.Getenv("DB_PORT"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Pwd, dbConfig.Host, dbConfig.Port, dbConfig.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	//db.AutoMigrate(&model.Borrower{})

	return db
}
