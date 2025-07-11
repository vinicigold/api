package database

import (
	"api/config"
	"api/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func Connect() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_USER"),
		config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect db.", err)
		os.Exit(2)
	}

	log.Println("running migrations")
	Db.AutoMigrate(&model.User{})
	Db.AutoMigrate(&model.Admin{})
	Db.AutoMigrate(&model.Employee{})
}
