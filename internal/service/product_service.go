package service

import "awesomeProject1/internal/models"

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetById(id int) (*models.Product, error)
	Create(product *models.Product) error
	Update(id int, student *models.Product) error
	Delete(id int) error
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(productRepo ProductRepository) *ProductService {
	return &ProductService{repo: productRepo}
}

func (p *ProductService) GetAllProducts() ([]models.Product, error) {
	return p.repo.GetAll()
}

func (p *ProductService) GetProductById(id int) (*models.Product, error) {
	return p.repo.GetById(id)
}

func (p *ProductService) CreateProduct(product *models.Product) error {
	return p.repo.Create(product)
}

func (p *ProductService) UpdateProduct(id int, product *models.Product) error {
	return p.repo.Update(id, product)
}

func (p *ProductService) DeleteProduct(id int) error {
	return p.repo.Delete(id)
}
