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
	ID            uuid.UUID       `json:"transaction_id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID        uuid.UUID       `json:"user_id"`
	Type          string          `json:"type"`
	Amount        decimal.Decimal `json:"amount"`
	BalanceBefore decimal.Decimal `json:"balance_before"`
	BalanceAfter  decimal.Decimal `json:"balance_after"`
	Remarks       string          `json:"remarks"`
	Status        string          `json:"status"`
	CreatedAt     time.Time       `json:"created_date"`
	UpdatedAt     time.Time       `json:"-"`
	DeletedAt     *time.Time      `json:"-"`
}

func (Transaction) TableName() string {
	return "transactions"
}
