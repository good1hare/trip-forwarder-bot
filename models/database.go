package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=0.0.0.0 user=postgres password=password dbname=trip_forwarder port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect PostgresSQL")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed migrate user")
	}
	return db
}
