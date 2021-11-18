package repository

import (
	"errors"
	"gorm.io/gorm"
	"superapp/models"
	"superapp/models/dto"
)

type ITransactionRepo interface {
	Transfer(trx *models.Transaction) (*models.Transaction, error)
	HistoryTrx(id, page, size string) ([]*dto.HistoryTrx, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func (t *transactionRepo) HistoryTrx(id, page, size string) ([]*dto.HistoryTrx, error) {
	offset, pageSize := Paginate(page, size)
	var err error
	histories := make([]*dto.HistoryTrx,0)
	if err = t.db.Debug().Raw("select trx.id, trx.transaction_date, trx.total_transaction, trx.status, ms.name as merchant_name from transactions trx join merchants ms on trx.merchant_id=ms.id where customer_id=? offset ? limit ?",id , offset, pageSize).Scan(&histories).Error ; err != nil{
		return nil, err
	}
	return histories, nil
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


