package repository

import (
	"awesomeProject1/internal/models"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	database *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{database: database}
}

func (p ProductRepositoryImpl) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := p.database.Find(&products).Error
	return products, err
}

func (p ProductRepositoryImpl) GetById(id int) (*models.Product, error) {
	var product models.Product
	err := p.database.First(&product, id).Error
	return &product, err
}

func (p ProductRepositoryImpl) Create(product *models.Product) error {
	return p.database.Create(product).Error
}

func (p ProductRepositoryImpl) Update(id int, product *models.Product) error {
	return p.database.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func (p ProductRepositoryImpl) Delete(id int) error {
	return p.database.Delete(&models.Product{}, id).Error
}
