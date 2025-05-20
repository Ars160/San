package main

import (
	"github.com/joho/godotenv"
	"log"

	"awesomeProject1/internal/database"
	"awesomeProject1/internal/delivery"
	"awesomeProject1/internal/repository"
	"awesomeProject1/internal/routes"
	"awesomeProject1/internal/service"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла")
	}

	db := database.InitDB()

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := controllers.NewProductHandler(productService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := controllers.NewCategoryHandler(categoryService)

	r := routes.SetupRoutes(productHandler, categoryHandler)

	log.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")
}
