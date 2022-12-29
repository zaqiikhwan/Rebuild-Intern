package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/domain"
)

var db *gorm.DB

func InitDB() error {
	// postgres://database_petlink_user:dV9cwr1wcVEhH3KsMf6QdSbHa4g2olSH@dpg-cemlf782i3molpj8srrg-a.singapore-postgres.render.com/database_petlink
	// "user=postgres password=Supadata1003/! host=db.xnhdcwxcwbbyexwldjzj.supabase.co TimeZone=Asia/Singapore port=5432 dbname=postgres"
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
