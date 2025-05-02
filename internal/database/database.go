package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// Получаем параметры подключения из переменных окружения
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Неверный порт: %v", err)
	}

	// Формируем строку подключения
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port,
	)

	// Создаём соединение с базой данных на уровне sql.DB
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	// Настраиваем драйвер миграций для PostgreSQL
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("Ошибка создания драйвера миграций:", err)
	}

	// Создаём экземпляр мигратора
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Ошибка инициализации мигратора:", err)
	}

	// Проверка на грязное состояние
	if version, dirty, err := m.Version(); err == nil && dirty {
		log.Printf("Обнаружено грязное состояние версии %d, исправляем...", version)
		if err := m.Force(int(version)); err != nil {
			log.Fatal("Ошибка исправления грязного состояния:", err)
		}
	}

	// Применяем все pending миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Ошибка применения миграций:", err)
	}

	// Инициализируем GORM с существующим соединением
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка инициализации GORM:", err)
	}

	log.Println("База данных подключена и миграции применены!")

	DB = db
	return db
}
