package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/domain"
)

var db *gorm.DB

func InitDB() error {
	_db, err := gorm.Open(postgres.Open("user=postgres password=Supadata1003/! host=db.xnhdcwxcwbbyexwldjzj.supabase.co TimeZone=Asia/Singapore port=5432 dbname=postgres"), &gorm.Config{})
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
