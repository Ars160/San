package controllers

import (
	"awesomeProject1/internal/middleware"
	"awesomeProject1/internal/service"
	"net/http"
	"strconv"

	"awesomeProject1/internal/models"
	"github.com/gin-gonic/gin"
)

func NewOutfitHandler(service *service.OutfitService) *OutfitHandler {
	return &OutfitHandler{service: service}
}

type OutfitHandler struct {
	service *service.OutfitService
}

func (h *OutfitHandler) GetAllOutfits(c *gin.Context) {
	outfits, _ := h.service.GetAllOutfits()
	c.JSON(http.StatusOK, outfits)
}

func (h *OutfitHandler) GetOutfitByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	outfit, err := h.service.GetOutfitById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Образ не найден"})
		return
	}
	c.JSON(http.StatusOK, outfit)
}

func (h *OutfitHandler) CreateOutfit(c *gin.Context) {
	userID, _, ok := middleware.GetUserFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется авторизация"})
		return
	}

	var outfit models.Outfit
	if err := c.ShouldBindJSON(&outfit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Устанавливаем владельца продукта
	outfit.UserID = userID

	err := h.service.CreateOutfit(&outfit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании образа"})
		return
	}

	c.JSON(http.StatusCreated, outfit)
}

func (h *OutfitHandler) UpdateOutfit(c *gin.Context) {
	currentUserID, _, ok := middleware.GetUserFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется авторизация"})
		return
	}

	// Получаем продукт из БД
	outfitID, _ := strconv.Atoi(c.Param("id"))
	existingOutfit, err := h.service.GetOutfitById(outfitID)

	// Проверяем, что пользователь - владелец
	if existingOutfit.UserID != currentUserID {
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
	existing, err := h.service.GetOutfitById(id)
	if err != nil || existing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Образ не найден"})
		return
	}

	var outfit models.Outfit
	if err := c.ShouldBindJSON(&outfit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}

	if err := h.service.UpdateOutfit(id, &outfit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении образа"})
		return
	}

	c.JSON(http.StatusOK, outfit)
}

func (h *OutfitHandler) DeleteOutfit(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	// Проверка, существует ли продукт
	existing, err := h.service.GetOutfitById(id)
	if err != nil || existing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Образ не найден"})
		return
	}

	if err := h.service.DeleteOutfit(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении образа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Образ успешно удалён"})
}
