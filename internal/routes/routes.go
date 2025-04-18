package routes

import (
	"awesomeProject1/internal/auth"
	"awesomeProject1/internal/delivery"
	"awesomeProject1/internal/middleware"
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
	}

	// Users
	users := r.Group("/auth")
	{
		users.POST("/login", auth.Login)
		users.POST("/register", auth.Register)
	}

	//Зашищенные
	protected := r.Group("/api")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/profile", auth.Profile)
		protected.POST("/products", controllers.CreateProduct)
		protected.DELETE("products/:id", controllers.DeleteProduct)
		protected.PUT("/:id", controllers.UpdateProduct)
	}

	return r
}
