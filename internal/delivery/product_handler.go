package controllers

import (
	"awesomeProject1/internal/middleware"
	"awesomeProject1/internal/service"
	"net/http"
	"strconv"

	"awesomeProject1/internal/models"
	"github.com/gin-gonic/gin"
)

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

type ProductHandler struct {
	service *service.ProductService
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, _ := h.service.GetAllProducts()
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	product, err := h.service.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	userID, _, ok := middleware.GetUserFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется авторизация"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Устанавливаем владельца продукта
	product.UserID = userID

	err := h.service.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании товара"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	currentUserID, _, ok := middleware.GetUserFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется авторизация"})
		return
	}

	// Получаем продукт из БД
	productID, _ := strconv.Atoi(c.Param("id"))
	existingProduct, err := h.service.GetProductById(productID)

	// Проверяем, что пользователь - владелец
	if existingProduct.UserID != currentUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Вы не можете изменять этот товар"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	// Проверяем — существует ли продукт
	existing, err := h.service.GetProductById(id)
	if err != nil || existing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}

	if err := h.service.UpdateProduct(id, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении товара"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	// Проверка, существует ли продукт
	existing, err := h.service.GetProductById(id)
	if err != nil || existing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	if err := h.service.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении товара"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно удалён"})
}
