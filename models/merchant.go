package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Merchant struct {
	ID 			uuid.UUID	`gorm:"type:uuid;unique;index" json:"id"`
	Name 		string		`gorm:"column:name" json:"name"`
	CreatedAt   time.Time  	`gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  	`gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time 	`gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (m *Merchant) Prepare()  {
	m.ID = uuid.NewV4()
	m.CreatedAt = time.Now()
}