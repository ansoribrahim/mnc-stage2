package data

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"user_id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName   string
	LastName    string
	PhoneNumber string
	Pin         string
	Address     string
	CreatedAt   time.Time  `json:"created_date"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}
