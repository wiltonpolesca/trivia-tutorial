package database

import (
	"fmt"
	"log"
	"os"

	"github.com/wiltonpolesca/go-trivia-tutorial/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=host.docker.internal user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{})

	DB = DbInstance{
		Db: db,
	}
}
