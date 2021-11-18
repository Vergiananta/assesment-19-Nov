package models

import (
	"errors"
	"github.com/badoux/checkmail"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type Customer struct {
	ID 			uuid.UUID 	`gorm:"type:uuid;unique;index " json:"id"`
	Name 		string 		`gorm:"not null; column:name" json:"name"`
	Address 	string 		`gorm:"not null: column:address " json:"address"`
	Email 		string		`gorm:"not null: column:email" json:"email"`
	Password 	string 		`gorm:"not null: column:password" json:"password"`
	IsActive 	bool		`gorm:" column:is_active" json:"is_active"`
	CreatedAt   time.Time  	`gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  	`gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time 	`gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (c *Customer) Prepare() error {
	c.ID = uuid.NewV4()
	hashedPassword, err := Hash(c.Password)
	if err != nil {
		return err
	}
	c.Password = string(hashedPassword)
	c.Email = html.EscapeString(strings.TrimSpace(c.Email))
	c.IsActive = true
	c.CreatedAt = time.Now()
	return nil
}

func (c *Customer) EditCustomer() error {
	hashedPassword, err := Hash(c.Password)
	if err != nil {
		return err
	}
	c.Password = string(hashedPassword)
	c.Email = html.EscapeString(strings.TrimSpace(c.Email))
	c.UpdatedAt = time.Now()
	return nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (c *Customer) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if c.Name == "" {
			return errors.New("Required Name")
		}
		if c.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return errors.New("Invalid email")
		}
		if c.Address == "" {
			return errors.New("Required Address")
		}
		if c.Password == "" {
			return errors.New("Required Password")
		}
		return nil

	default:
		if c.Name == "" {
			return errors.New("Required Name")
		}
		if c.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return errors.New("Invalid email")
		}
		if c.Address == "" {
			return errors.New("Required Address")
		}
		if c.Password == "" {
			return errors.New("Required Password")
		}

		return nil
	}

}