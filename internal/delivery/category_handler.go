package controllers

import (
	"awesomeProject1/internal/service"
	"net/http"
	"strconv"

	"awesomeProject1/internal/models"
	"github.com/gin-gonic/gin"
)

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

type CategoryHandler struct {
	service *service.CategoryService
}

func (h *CategoryHandler) GetAllCategory(c *gin.Context) {
	category, _ := h.service.GetAllCategory()
	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	category, err := h.service.GetCategoryById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найден"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}

	err := h.service.CreateCategory(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании категорий"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	// Проверяем — существует ли продукт
	existing, err := h.service.GetCategoryById(id)
	if err != nil || existing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найден"})
		return
	}

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}

	if err := h.service.UpdateCategory(id, &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении категорий"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	// Проверка, существует ли продукт
	existing, err := h.service.GetCategoryById(id)
	if err != nil || existing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найден"})
		return
	}

	if err := h.service.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении категорий"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Категория успешно удалён"})
}
