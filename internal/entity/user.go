package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User represents resources in JSON format.
type User struct {
	ID           primitive.ObjectID `bson:"_id, omitempty" json:"_id"`
	Login        string             `json:"login" `
	PasswordHash string             `json:"-" `
	Email        string             `json:"-" `
	Website      string             `json:"website"`
	Token        string             `json:"-" `
	Banned       bool               `json:"banned" `
	CreatedAt    time.Time          `json:"created_at" `
	UpdatedAt    time.Time          `json:"updated_at"`
}
