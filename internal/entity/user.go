package entity

import (
	"time"
)

//User represents resources in JSON format.
type User struct {
	ID        int       `json:"id"`
	Password  string    `json:"-" `
	Login     string    `json:"login" `
	Email     string    `json:"email,omitempty" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//GetPublicProfile - returns only public information
func (u User) GetPublicProfile() User {
	u.Email = ""
	u.Password = ""
	return u
}

//TableName - returns table name in database
func (u User) TableName() string {
	return "users"
}
