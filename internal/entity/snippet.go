package entity

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

//Snippet - ...
type Snippet struct {
	ID            int       `from:"id"`
	UserID        int       `from:"user_id"`
	Favorite      bool      `from:"favorite"`
	Access        int       `from:"access"`
	Title         string    `from:"title"`
	Content       null.String    `from:"content"`
	FileExtension string    `from:"file_extension"`
	EditorOptions struct{}  `from:"editor_options"`
	CreatedAt     time.Time `from:"created_at"`
	UpdatedAt     time.Time `from:"updated_at"`
}

//TableName - returns table name in database
func (s Snippet) TableName() string {
	return "snippets"
}
