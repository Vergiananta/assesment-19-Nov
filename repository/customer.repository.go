package repository

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"superapp/middlewares"
	"superapp/models"
	"superapp/models/dto"
	"time"
)

type ICustomerRepo interface {
	CreateCustomer(newCust *models.Customer) (*models.Customer, error)
	UpdateCustomer(newCust *models.Customer) (*dto.CustomerResponse, error)
	FindByIdCustomer(id string) (*models.Customer, error)
	DeleteCustomer(id string) (string, error)
	LoginCustomer(cust *dto.LoginRequest) (string, error)
}

type customerRepo struct {
	db *gorm.DB
}

func (c *customerRepo) LoginCustomer(cust *dto.LoginRequest) (string, error) {
	var err error
	var accounts models.Customer

	if err = c.db.Debug().Table("customers").Where("email = ?", cust.Email).Find(&accounts).Error; err != nil {
		return "", err
	}
	err = models.VerifyPassword(accounts.Password, cust.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	if gorm.ErrRecordNotFound == err || accounts.Name == "" || accounts.IsActive == false {
		return "", errors.New("User Not Found")
	}
	return middlewares.CreateToken(&accounts)
}

func (c *customerRepo) CreateCustomer(newCust *models.Customer) (*models.Customer, error) {
	var err error
	fmt.Println(newCust)
	if err = c.db.Debug().Create(&newCust).Error; err != nil {
		return nil, err
	}
	return newCust, nil
}

func (c *customerRepo) UpdateCustomer(newCust *models.Customer) (*dto.CustomerResponse, error) {

	if err := c.db.Debug().Save(newCust).Error; err != nil {
		return nil, err
	}
	return &dto.CustomerResponse{
		Name: newCust.Name,
		Address: newCust.Address,
		Email: newCust.Email,
	}, nil
}

func (c *customerRepo) FindByIdCustomer(id string) (*models.Customer, error) {
	var err error
	uid, _ := uuid.FromString(id)
	var user models.Customer
	err = c.db.Debug().Table("customers").Where("id  = ?", uid).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.ErrRecordNotFound == err {
		return nil, errors.New("User Not Found")
	}
	return &user, nil
}

func (c *customerRepo) DeleteCustomer(id string) (string, error) {
	uid, _ := uuid.FromString(id)
	var account models.Customer
	c.db = c.db.Debug().Model(&account).Where("id = ?", uid).Take(&account).UpdateColumns(
		map[string]interface{}{
			"is_active":  false,
			"deleted_at": time.Now(),
		},
	)
	if err := c.db.Error; err != nil {
		return "",errors.New("Customer Not Found")
	}
	return "Data has been deleted",nil
}

func NewCustomerRepo(db *gorm.DB) ICustomerRepo  {
	return &customerRepo{db,}
}
