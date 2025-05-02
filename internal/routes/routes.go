package routes

import (
	"awesomeProject1/internal/auth"
	"awesomeProject1/internal/delivery"
	"awesomeProject1/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(productHandler *controllers.ProductHandler) *gin.Engine {
	r := gin.Default()

	// Группа /products
	products := r.Group("/products")
	{
		products.GET("/", productHandler.GetAllProducts)
		products.GET("/:id", productHandler.GetProductByID)
	}

	// Users
	users := r.Group("/auth")
	{
		users.POST("/login", auth.Login)
		users.POST("/register", auth.Register)
	}

	// Защищенные
	protected := r.Group("/api")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/profile", auth.Profile)

		// Admin-only routes
		admin := protected.Group("/products")
		admin.Use(middleware.RequireRole("admin"))
		{
			admin.POST("/", productHandler.CreateProduct)
			admin.PUT("/:id", productHandler.UpdateProduct)
			admin.DELETE("/:id", productHandler.DeleteProduct)
		}
	}

	return r
}
