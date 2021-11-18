package usecase

import (
	"superapp/models"
	"superapp/models/dto"
	"superapp/repository"
)

type ICustomerUsecase interface {
	CreateCustomer(cust *dto.CustomerRequest) (*dto.CustomerResponse, error)
	UpdateCustomer(cust *dto.CustomerRequest, id string) (*dto.CustomerResponse, error)
	LoginCustomer(cust *dto.LoginRequest) (string, error)
	DeleteCustomer(id string) (string, error)
	FindByIdCustomer(id string) (*models.Customer, error)
}

type customerUsecase struct {
	customerRepo repository.ICustomerRepo
}

func (c *customerUsecase) FindByIdCustomer(id string) (*models.Customer, error) {
	return c.customerRepo.FindByIdCustomer(id)
}

func (c *customerUsecase) CreateCustomer(cust *dto.CustomerRequest) (*dto.CustomerResponse, error) {

	customer := models.Customer{
		Name: cust.Name,
		Address: cust.Address,
		Email: cust.Email,
		Password: cust.Password,
	}
	customer.Prepare()
	customer.Validate("")

	c.customerRepo.CreateCustomer(&customer)

	return &dto.CustomerResponse{
		Name: customer.Name,
		Address: customer.Address,
		Email: customer.Email,
	}, nil
}

func (c *customerUsecase) UpdateCustomer(cust *dto.CustomerRequest, id string) (*dto.CustomerResponse, error) {
	temp, err := c.FindByIdCustomer(id)
	temp.Name = cust.Name
	temp.Email = cust.Email
	temp.Password = cust.Password
	temp.Address = cust.Address
	temp.Validate("update")
	temp.EditCustomer()
	if err != nil {
		return nil, err
	}
	return c.customerRepo.UpdateCustomer(temp)
}

func (c *customerUsecase) LoginCustomer(cust *dto.LoginRequest) (string, error) {
	return c.customerRepo.LoginCustomer(cust)
}

func (c *customerUsecase) DeleteCustomer(id string) (string, error) {
	return c.customerRepo.DeleteCustomer(id)
}

func NewCustomerUsecase(customerRepo repository.ICustomerRepo) ICustomerUsecase  {
	return &customerUsecase{
		customerRepo,
	}
}