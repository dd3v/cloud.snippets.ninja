package entity

import "time"

//Session represents refresh tokens
type Session struct {
	ID           int
	UserID       int
	RefreshToken string
	Exp          time.Time
	IP           string
	UserAgent    string
	CreatedAt    time.Time
}

func (s Session) TableName() string {
	return "sessions"
}
