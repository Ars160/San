package service

import "awesomeProject1/internal/models"

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetById(id int) (*models.Category, error)
	Create(category *models.Category) error
	Update(id int, category *models.Category) error
	Delete(id int) error
}

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(categoryRepo CategoryRepository) *CategoryService {
	return &CategoryService{repo: categoryRepo}
}

func (c *CategoryService) GetAllCategory() ([]models.Category, error) {
	return c.repo.GetAll()
}

func (c *CategoryService) GetCategoryById(id int) (*models.Category, error) {
	return c.repo.GetById(id)
}

func (c *CategoryService) CreateCategory(category *models.Category) error {
	return c.repo.Create(category)
}

func (c *CategoryService) UpdateCategory(id int, category *models.Category) error {
	return c.repo.Update(id, category)
}

func (c *CategoryService) DeleteCategory(id int) error {
	return c.repo.Delete(id)
}
