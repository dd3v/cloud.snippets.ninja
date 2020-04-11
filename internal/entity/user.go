package entity

import (
	"time"
)

//User represents resources in JSON format.
type User struct {
	ID           int       `json:"id"`
	PasswordHash string    `json:"-" `
	Login        string    `json:"login" `
	Email        string    `json:"-" `
	Website      string    `json:"website"`
	Banned       bool      `json:"banned" `
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u User) TableName() string {
	return "users"
}
