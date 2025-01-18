package data

import "time"

type User struct {
	UserID      string    `json:"user_id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number" gorm:"uniqueIndex"`
	Address     string    `json:"address"`
	CreatedDate time.Time `json:"created_date" gorm:"autoCreateTime"`
}
