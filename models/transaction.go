package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type statusTransaction string

const (
	TRANSFER statusTransaction = "TRANSFER"
	TOPUP 	 statusTransaction = "TOPUP"
	WITHDRAW statusTransaction = "WITHDRAW"
	)
type Transaction struct {
	ID 					uuid.UUID			`gorm:"type:uuid;unique;index" json:"id"`
	TransactionDate 	string				`gorm:"column: transaction_date" json:"transaction_date"`
	TotalTransaction	int64  				`gorm:"column: total_transaction" json:"total_transaction"`
	Status				statusTransaction 	`gorm:"column:status" json:"status"`
	MerchantID 			uuid.UUID			`gorm:"null" json:"merchant_id"`
	Merchant	  		*Merchant			`gorm:"foreignKey:MerchantID;references:ID;not null" json:"merchant"`
	CustomerID 			uuid.UUID			`gorm:"null" json:"customer_id"`
	Customer 			*Customer			`gorm:"foreignKey:CustomerID;references:ID;not null" json:"customer"`
	CreatedAt   		time.Time  			`gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   		time.Time  			`gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   		*time.Time 			`gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (t *Transaction) Prepare() error{
	t.ID = uuid.NewV4()
	t.CreatedAt = time.Now()
	return nil
}
