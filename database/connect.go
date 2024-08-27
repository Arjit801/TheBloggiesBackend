package database

import (
	"log"
	"os"

	"github.com/Arjit801/TheBloggies/dao/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect() {
	// databaseURL := "postgres://postgres:987654321@localhost:5432/TheBloggies"
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading .env file")
	}
	dsn := os.Getenv("DSN")
	// Create a connection pool
	// conn, err := pgxpool.New(context.Background(), databaseURL)/
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	} else {
		log.Println("connect successfully")
	}
	// Check the connection
	DB = db
	db.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}