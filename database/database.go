package database

import (
	"fmt",
	"log",
	"os",
	"github.com/jacob1225/instagram-bot/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type Database struct {
	Db *gorm.Db

}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Toronto Canada/Eastern",
		os.GetEnv("DB_USER"), 
		os.GetEnv("DB_PASSWORD"),
		os.GetEnv("DB_NAME")
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected!")
	db.logger = logger.Default.LogMode(Logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.User{})

	DB = Dbinstance{
		Db: db,
	}
}