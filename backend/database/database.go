package database

import (
	// "database/sql"
	"log"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"backend/models"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// ConnectDB connects to the postgres database
func ConnectDB() {
	dotenverr := godotenv.Load(".env")
	
	if dotenverr != nil {
		panic("Failed to load .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

    db.AutoMigrate(&models.Organization{}, &models.User{}, &models.Event{}, &models.EventType{})

	DB = Dbinstance{
		Db: db,
	}
}