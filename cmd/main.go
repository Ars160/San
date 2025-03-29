package main

import (
	"log"

	"awesomeProject1/internal/database"
	"awesomeProject1/internal/routes"
)

func main() {
	database.InitDB()

	r := routes.SetupRoutes()

	log.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")
}
