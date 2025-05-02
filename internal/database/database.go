package database

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := "localhost"
	port := 5430
	user := "admin"
	password := "test123"
	dbname := "san_db"

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

	// Применяем миграции
	if version, dirty, err := m.Version(); err == nil && dirty {
		log.Printf("Обнаружено грязное состояние версии %d, исправляем...", version)
		if err := m.Force(int(version)); err != nil {
			log.Fatal("Ошибка исправления грязного состояния:", err)
		}
	}

	if err != nil {
		log.Fatal("Ошибка инициализации мигратора:", err)
	}

	// Применяем все pending миграции
	if err := m.Up(); err != migrate.ErrNoChange && err != nil {
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

	// Возвращаем соединение
	return db
}
