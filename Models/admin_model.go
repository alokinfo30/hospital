// admin_model.go

package model

import "time"

// Admin represents the Admin model
type Admin struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewAdmin creates a new Admin instance
func NewAdmin(name, email, phone, password string) *Admin {
	return &Admin{
		Name:      name,
		Email:     email,
		Phone:     phone,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
