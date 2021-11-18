package usecase

import (
	"superapp/models"
	"superapp/models/dto"
	"superapp/repository"
)

type ITransactionUsecase interface {
	Transfer(trx *models.Transaction) (*models.Transaction, error)
	HistoryTrx(id, page, size string) ([]*dto.HistoryTrx, error)
}

type transactionUsecase struct {
	transactionRepo 	repository.ITransactionRepo
}

func (t *transactionUsecase) HistoryTrx(id, page, size string) ([]*dto.HistoryTrx, error) {
	return t.transactionRepo.HistoryTrx(id, page, size)
}

func (t *transactionUsecase) Transfer(trx *models.Transaction) (*models.Transaction, error) {
	trx.Prepare()
	trx.Status = models.TRANSFER
	return t.transactionRepo.Transfer(trx)
}

func NewTransactionUsecase(trxRepo repository.ITransactionRepo) ITransactionUsecase {
	return &transactionUsecase{trxRepo}
}
