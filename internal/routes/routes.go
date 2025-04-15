package routes

import (
	"awesomeProject1/internal/auth"
	"awesomeProject1/internal/delivery"
	"github.com/gin-gonic/gin"
)

// SetupRoutes — создаём роутер и регистрируем эндпойнты
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Группа /products
	products := r.Group("/products")
	{
		products.GET("/", controllers.GetAllProducts)
		products.GET("/:id", controllers.GetProductByID)
		products.POST("/", controllers.CreateProduct)
		products.PUT("/:id", controllers.UpdateProduct)
		products.DELETE("/:id", controllers.DeleteProduct)
	}

	// Users
	users := r.Group("/auth")
	{
		users.POST("/login", auth.Login)
		users.POST("/register", auth.Register)
	}

	return r
}
