package manager

import (
	"superapp/connect"
	"superapp/usecase"
)

type ServiceManager interface {
	CustomerUsecase() usecase.ICustomerUsecase
	MerchantUsecase() usecase.IMerchantUsecase
	TransactionUsecase() usecase.ITransactionUsecase
}

type serviceManager struct {
	repo RepoManager
}

func (sm *serviceManager) TransactionUsecase() usecase.ITransactionUsecase {
	return usecase.NewTransactionUsecase(sm.repo.TransactionRepo())
}

func (sm *serviceManager) MerchantUsecase() usecase.IMerchantUsecase {
	return usecase.NewMerchantUsecase(sm.repo.MerchantRepo())
}

func (sm *serviceManager) CustomerUsecase() usecase.ICustomerUsecase {
	return usecase.NewCustomerUsecase(sm.repo.CustomerRepo())
}

func NewServiceManager(connect connect.Connect) ServiceManager {
	return &serviceManager{repo: NewRepoManager(connect)}
}
