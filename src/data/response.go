package data

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type RegisterResponse struct {
	Status  string        `json:"status"`
	Result  *UserResponse `json:"result,omitempty"`
	Message *string       `json:"message,omitempty"`
}

type UserResponse struct {
	ID          uuid.UUID `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_date"`
}

type LoginResponse struct {
	Status  string    `json:"status"`
	Result  *LoginRsp `json:"result,omitempty"`
	Message *string   `json:"message,omitempty"`
}

type LoginRsp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TopUpResponse struct {
	Status  string    `json:"status"`
	Result  *TopUpRsp `json:"result,omitempty"`
	Message *string   `json:"message,omitempty"`
}

type TopUpRsp struct {
	TopUpID       string          `json:"top_up_id"`
	AmountTopUp   decimal.Decimal `json:"amount_top_up"`
	BalanceBefore decimal.Decimal `json:"balance_before"`
	BalanceAfter  decimal.Decimal `json:"balance_after"`
	CreatedDate   string          `json:"created_date"`
}

type PaymentResponse struct {
	Status  string      `json:"status"`
	Result  *PaymentRsp `json:"result,omitempty"`
	Message *string     `json:"message,omitempty"`
}

type PaymentRsp struct {
	PaymentID     string          `json:"payment_id"`
	Amount        decimal.Decimal `json:"amount"`
	Remarks       string          `json:"remarks"`
	BalanceBefore decimal.Decimal `json:"balance_before"`
	BalanceAfter  decimal.Decimal `json:"balance_after"`
	CreatedDate   string          `json:"created_date"`
}

type TransferResponse struct {
	Status  string       `json:"status"`
	Result  *TransferRsp `json:"result,omitempty"`
	Message *string      `json:"message,omitempty"`
}

type TransferRsp struct {
	TransferID    string          `json:"transfer_id"`
	Amount        decimal.Decimal `json:"amount"`
	Remarks       string          `json:"remarks"`
	BalanceBefore decimal.Decimal `json:"balance_before"`
	BalanceAfter  decimal.Decimal `json:"balance_after"`
	CreatedDate   string          `json:"created_date"`
}

type TransactionResponse struct {
	Status  string        `json:"status"`
	Result  []Transaction `json:"result,omitempty"`
	Message *string       `json:"message,omitempty"`
}
