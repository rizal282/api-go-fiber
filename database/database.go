package database

import (
	"log"
	"os"

	"github.com/rizal282/go-fiber-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

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