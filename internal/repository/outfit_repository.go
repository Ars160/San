package repository

import (
	"awesomeProject1/internal/models"
	"gorm.io/gorm"
)

type OutfitRepositoryImpl struct {
	database *gorm.DB
}

func NewOutfitRepository(database *gorm.DB) *OutfitRepositoryImpl {
	return &OutfitRepositoryImpl{database: database}
}

func (r OutfitRepositoryImpl) GetAll() ([]models.Outfit, error) {
	var outfits []models.Outfit
	err := r.database.Preload("Products").Find(&outfits).Error
	return outfits, err
}

func (r OutfitRepositoryImpl) GetById(id int) (*models.Outfit, error) {
	var outfit models.Outfit
	err := r.database.
		Preload("Products"). // Загружаем связанные продукты
		First(&outfit, id).  // Ищем аутфит по ID
		Error

	return &outfit, err
}

func (r OutfitRepositoryImpl) Create(outfit *models.Outfit) error {
	return r.database.Create(outfit).Error
}

func (r OutfitRepositoryImpl) Update(id int, outfit *models.Outfit) error {
	return r.database.Model(&models.Outfit{}).Where("id = ?", id).Updates(outfit).Error
}

func (r OutfitRepositoryImpl) Delete(id int) error {
	return r.database.Delete(&models.Outfit{}, id).Error
}
