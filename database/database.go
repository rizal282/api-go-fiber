package database

import (
	"log"
	"os"

	"github.com/rizal282/go-fiber-api/models"
	"gorm.io/driver/mysql"
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {

	// DSN: "iniusername:inipassword@tcp(127.0.0.1:3306)/db_golang?charset=utf8&parseTime=True&loc=Local",

	// db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	// config for connect to mysql database

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root@tcp(127.0.0.1:3306)/db_golang?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{})

	

	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())

		os.Exit(2)
	}

	log.Println("Connected successfully to Database")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	// add Migrations here 
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}