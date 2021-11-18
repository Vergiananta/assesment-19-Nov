package manager

import (
	"superapp/connect"
	"superapp/repository"
)

type RepoManager interface {
	CustomerRepo() repository.ICustomerRepo
	MerchantRepo() repository.IMerchantRepo
	TransactionRepo() repository.ITransactionRepo
}

type repoManager struct {
	connect connect.Connect
}

func (rm *repoManager) TransactionRepo() repository.ITransactionRepo {
	return repository.NewTransactionRepo(rm.connect.SqlDb())
}

func (rm *repoManager) MerchantRepo() repository.IMerchantRepo {
	return repository.NewMerchantRepo(rm.connect.SqlDb())
}

func (rm *repoManager) CustomerRepo() repository.ICustomerRepo {
	return repository.NewCustomerRepo(rm.connect.SqlDb())
}

func NewRepoManager(connect connect.Connect) RepoManager {
	return &repoManager{connect}
}

