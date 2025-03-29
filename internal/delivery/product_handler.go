package controllers

import (
	"net/http"
	"strconv"

	"awesomeProject1/internal/database"
	"awesomeProject1/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Product

	// SELECT * FROM products
	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения из БД"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductByID — получить один товар по ID
func GetProductByID(c *gin.Context) {
	idParam := c.Param("id") // /products/10 => id = "10"
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var product models.Product
	// SELECT * FROM products WHERE id = ?
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct — создать новый товар
func CreateProduct(c *gin.Context) {
	var product models.Product

	// Парсим JSON из тела запроса
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}

	// INSERT INTO products (...)
	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании товара"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct — обновить товар
func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var product models.Product
	// Ищем товар
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	// Считываем новые данные
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}

	// Обновляем поля
	product.Name = input.Name
	product.Price = input.Price
	product.Category = input.Category
	product.Stock = input.Stock

	// Сохраняем
	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении товара"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct — удалить товар
func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	// DELETE FROM products WHERE id = ?
	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении товара"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно удален"})
}
