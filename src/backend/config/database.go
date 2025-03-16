package config

import (
	"log"
	"backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("gpt.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke SQLite: %v", err)
	}

	log.Println("Database SQLite berhasil terhubung!")

	DB = database

	err = DB.AutoMigrate(&models.ChatGPT{})
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}
}