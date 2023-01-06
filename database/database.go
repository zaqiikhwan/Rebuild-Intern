package database

import (
	"fmt"
	_"log"
	"os"

	_"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/domain"
)

var db *gorm.DB

func InitDB() error {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// dsn := fmt.Sprintf(
	// 	"user=%s " + 
	// 	"password=%s " + 
	// 	"host=%s " + 
	// 	"port=%s " +
	// 	"dbname=%s ",
	// os.Getenv("RENDER_USER"),
	// os.Getenv("RENDER_PASSWORD"),
	// os.Getenv("RENDER_HOST"),
	// os.Getenv("RENDER_PORT"),
	// os.Getenv("RENDER_DBNAME"))

	// how connect database using render (PostgreSQL)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s", os.Getenv("RENDER_USER"), os.Getenv("RENDER_PASSWORD"), os.Getenv("RENDER_HOST"), os.Getenv("RENDER_DBNAME"))
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = _db
	err = db.AutoMigrate(&domain.Hospital{}, &domain.User{}, &domain.Doctor{})
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
