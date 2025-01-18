package data

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"user_id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName   string
	LastName    string
	PhoneNumber string
	Pin         string
	Address     string
	Balance     decimal.Decimal
	CreatedAt   time.Time  `json:"created_date"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

func (User) TableName() string {
	return "users"
}

type Transaction struct {
	ID            uuid.UUID `json:"transaction_id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID        uuid.UUID
	Type          string
	Amount        decimal.Decimal
	BalanceBefore decimal.Decimal
	BalanceAfter  decimal.Decimal
	Remarks       string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

func (Transaction) TableName() string {
	return "transactions"
}
