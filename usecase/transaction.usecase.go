package usecase

import (
	"superapp/models"
	"superapp/repository"
)

type ITransactionUsecase interface {
	Transfer(trx *models.Transaction) (*models.Transaction, error)
}

type transactionUsecase struct {
	transactionRepo 	repository.ITransactionRepo
}

func (t *transactionUsecase) Transfer(trx *models.Transaction) (*models.Transaction, error) {
	trx.Prepare()
	trx.Status = models.TRANSFER
	return t.transactionRepo.Transfer(trx)
}

func NewTransactionUsecase(trxRepo repository.ITransactionRepo) ITransactionUsecase {
	return &transactionUsecase{trxRepo}
}
