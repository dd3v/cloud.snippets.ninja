package entity

import "time"

//Snippet - ...
type Snippet struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Favorite      bool      `json:"favorite"`
	Public        bool      `json:"public"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	FileExtension string    `json:"file_extension"`
	EditorOptions struct{}  `json:"editor_options"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

//TableName - returns table name in database
func (s Snippet) TableName() string {
	return "snippets"
}
