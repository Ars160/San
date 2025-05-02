package main

import (
	"log"

	"awesomeProject1/internal/database"
	"awesomeProject1/internal/delivery"
	"awesomeProject1/internal/repository"
	"awesomeProject1/internal/routes"
	"awesomeProject1/internal/service"
)

func main() {
	db := database.InitDB()

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := controllers.NewProductHandler(productService)

	r := routes.SetupRoutes(productHandler)

	log.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")
}
