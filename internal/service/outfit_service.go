package service

import "awesomeProject1/internal/models"

type OutfitRepository interface {
	GetAll() ([]models.Outfit, error)
	GetById(id int) (*models.Outfit, error)
	Create(outfit *models.Outfit) error
	Update(id int, outfit *models.Outfit) error
	Delete(id int) error
}

type OutfitService struct {
	repo OutfitRepository
}

func NewOutfitService(outfitRepo OutfitRepository) *OutfitService {
	return &OutfitService{repo: outfitRepo}
}

func (o *OutfitService) GetAllOutfits() ([]models.Outfit, error) {
	return o.repo.GetAll()
}

func (o *OutfitService) GetOutfitById(id int) (*models.Outfit, error) {
	return o.repo.GetById(id)
}

func (o *OutfitService) CreateOutfit(outfit *models.Outfit) error {
	return o.repo.Create(outfit)
}

func (o *OutfitService) UpdateOutfit(id int, outfit *models.Outfit) error {
	return o.repo.Update(id, outfit)
}

func (o *OutfitService) DeleteOutfit(id int) error {
	return o.repo.Delete(id)
}
