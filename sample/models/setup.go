package models

import (
	"log"

	"gorm.io/driver/sqlite" // O driver que você quer usar
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Conectando ao SQLite usando o driver sem CGO
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	// Migra a estrutura de 'Book' para o banco de dados
	database.AutoMigrate(&Book{})

	DB = database
}
