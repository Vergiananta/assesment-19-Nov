package repository

import (
	"errors"
	"gorm.io/gorm"
	"superapp/models"
)

type ITransactionRepo interface {
	Transfer(trx *models.Transaction) (*models.Transaction, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func (t *transactionRepo) Transfer(trx *models.Transaction) (*models.Transaction, error) {
	if err := t.db.Debug().Create(trx).Error; err != nil {
		return nil, errors.New("Something wrong input")
	}
	return trx, nil
}

func NewTransactionRepo(db *gorm.DB) ITransactionRepo {
	return &transactionRepo{db}
}


