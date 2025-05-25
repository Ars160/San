package routes

import (
	"awesomeProject1/internal/auth"
	"awesomeProject1/internal/delivery"
	"awesomeProject1/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(productHandler *controllers.ProductHandler, categoryHandler *controllers.CategoryHandler, outfitHandler *controllers.OutfitHandler) *gin.Engine {
	r := gin.Default()

	//category /category
	category := r.Group("/category")
	{
		category.GET("/", categoryHandler.GetAllCategory)
		category.GET("/:id", categoryHandler.GetCategoryByID)
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

		// Products (под /api/products)
		products := protected.Group("/products")
		{
			products.GET("/", productHandler.GetAllProducts)
			products.GET("/:id", productHandler.GetProductByID)
			products.POST("/", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}

		outfits := protected.Group("/outfits")
		{
			outfits.GET("/", outfitHandler.GetAllOutfits)
			outfits.GET("/:id", outfitHandler.GetOutfitByID)
			outfits.POST("/", outfitHandler.CreateOutfit)
			outfits.PUT("/:id", outfitHandler.UpdateOutfit)
			outfits.DELETE("/:id", outfitHandler.DeleteOutfit)
		}

		// Admin-only routes
		admin := protected.Group("/category")
		admin.Use(middleware.RequireRole("admin"))
		{
			admin.POST("/", categoryHandler.CreateCategory)
			admin.PUT("/:id", categoryHandler.UpdateCategory)
			admin.DELETE("/:id", categoryHandler.DeleteCategory)
		}
	}

	return r
}
