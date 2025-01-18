package data

import (
	"time"

	"github.com/google/uuid"
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
