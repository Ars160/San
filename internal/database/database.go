package database

import (
	"fmt"
	"log"
	_ "os"

	"awesomeProject1/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	// Настройки БД (можно заменить на переменные окружения)
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "123123"
	dbname := "san_db"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	// Автоматическая миграция
	DB.AutoMigrate(&models.Product{})
	log.Println("База данных подключена и миграции применены!")
}
