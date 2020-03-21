package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//RefreshTokens represents refresh tokens in JSON format.
type RefreshTokens struct {
	ID        primitive.ObjectID
	Token     string
	Exp       time.Time
	IP        string
	UserAgent string
	CreatedAt time.Time
}

//User represents resources in JSON format.
type User struct {
	ID            primitive.ObjectID `bson:"_id, omitempty" json:"_id"`
	PasswordHash  string             `bson:"passwordHash"  json:"-" `
	RefreshTokens []RefreshTokens    `bson:"refreshTokens" json:"-"`
	Login         string             `json:"login" `
	Email         string             `json:"-" `
	Website       string             `json:"website"`
	Token         string             `json:"-" `
	Banned        bool               `json:"banned" `
	CreatedAt     time.Time          `bson:"createdAt" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updatedAt" json:"updated_at"`
}
