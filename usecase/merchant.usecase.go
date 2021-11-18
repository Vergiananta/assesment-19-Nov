package usecase

import (
	"superapp/models"
	"superapp/repository"
)

type IMerchantUsecase interface {
	CreateMerchant(newMerchant *models.Merchant) (*models.Merchant, error)
	DeleteMerchant(id string) error
	UpdateMerchant(merchant *models.Merchant) (*models.Merchant, error)
	GetAllMerchant(page, size string) ([]*models.Merchant, error)
}

type merchantUsecase struct {
	merchantRepo repository.IMerchantRepo
}

func (m *merchantUsecase) UpdateMerchant(merchant *models.Merchant) (*models.Merchant, error) {
	return m.merchantRepo.UpdateMerchant(merchant)
}

func (m *merchantUsecase) GetAllMerchant(page, size string) ([]*models.Merchant, error) {
	return m.merchantRepo.FindAllMerchant(page,size)
}

func (m *merchantUsecase) CreateMerchant(newMerchant *models.Merchant) (*models.Merchant, error) {
	newMerchant.Prepare()
	return m.merchantRepo.CreateMerchant(newMerchant)
}

func (m *merchantUsecase) DeleteMerchant(id string) error {
	return m.merchantRepo.DeleteMerchant(id)
}

func NewMerchantUsecase(repo repository.IMerchantRepo) IMerchantUsecase  {
	return &merchantUsecase{repo}
}
