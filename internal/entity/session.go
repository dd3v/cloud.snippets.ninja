package entity

import "time"

//Session represents refresh tokens
type Session struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	Exp          time.Time `json:"exp"`
	IP           string    `json:"ip"`
	UserAgent    string    `json:"user_agent"`
	CreatedAt    time.Time `json:"created_at"`
}

//TokenPair represents token pars - access_token and refresh_token
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//TableName - returns sessions table name from database
func (s Session) TableName() string {
	return "sessions"
}
