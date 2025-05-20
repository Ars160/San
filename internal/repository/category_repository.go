package repository

import (
	"awesomeProject1/internal/models"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	database *gorm.DB
}

func NewCategoryRepository(database *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{database: database}
}

func (c CategoryRepositoryImpl) GetAll() ([]models.Category, error) {
	var category []models.Category
	err := c.database.Find(&category).Error
	return category, err
}

func (c CategoryRepositoryImpl) GetById(id int) (*models.Category, error) {
	var category models.Category
	err := c.database.First(&category, id).Error
	return &category, err
}

func (c CategoryRepositoryImpl) Create(category *models.Category) error {
	return c.database.Create(category).Error
}

func (c CategoryRepositoryImpl) Update(id int, category *models.Category) error {
	return c.database.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error
}

func (c CategoryRepositoryImpl) Delete(id int) error {
	return c.database.Delete(&models.Category{}, id).Error
}
