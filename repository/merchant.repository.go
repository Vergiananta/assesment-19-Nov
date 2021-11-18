package repository

import (
	"gorm.io/gorm"
	"superapp/models"
)

type IMerchantRepo interface {
	CreateMerchant(newMerchant *models.Merchant) (*models.Merchant, error)
	UpdateMerchant(merchant *models.Merchant) (*models.Merchant, error)
	DeleteMerchant(id string) error
	FindAllMerchant(page, size string) ([]*models.Merchant, error)
}

type merchantRepo struct {
	db 	*gorm.DB
}

func (m *merchantRepo) FindAllMerchant(page, size string) ([]*models.Merchant, error) {
	offset, pageSize := Paginate(page, size)
	var err error
	merchant := make([]*models.Merchant,0)
	if err = m.db.Debug().Raw("select * from merchants offset ? limit ?", offset,pageSize).Scan(&merchant).Error ; err != nil{
		return nil, err
	}
	return merchant, nil
}

func (m *merchantRepo) CreateMerchant(newMerchant *models.Merchant) (*models.Merchant, error) {
	var err error
	if err = m.db.Debug().Create(&newMerchant).Error ;err != nil {
		return nil, err
	}
	return newMerchant, nil
}

func (m *merchantRepo) UpdateMerchant(merchant *models.Merchant) (*models.Merchant, error) {
	if err:= m.db.Debug().Save(merchant).Error; err != nil {
		return nil, err
	}
	return merchant, nil
}

func (m *merchantRepo) DeleteMerchant(id string) error {
	m.db = m.db.Debug().Model(&models.Merchant{}).Where("id = ?", id).Take(&models.Merchant{}).Delete(&models.Merchant{})
	if m.db.Error != nil {
		return m.db.Error
	}
	return nil

}

func NewMerchantRepo(db *gorm.DB) IMerchantRepo {
	return &merchantRepo{
		db,
	}
}
