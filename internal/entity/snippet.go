package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gopkg.in/guregu/null.v4"
)

//Snippet - ...
type Snippet struct {
	ID                  int                 `json:"id"`
	UserID              int                 `json:"user_id"`
	Favorite            bool                `json:"favorite"`
	AccessLevel         int                 `json:"access_level"`
	Title               string              `json:"title"`
	Content             null.String         `json:"content"`
	Language            string              `json:"language"`
	CustomEditorOptions CustomEditorOptions `json:"custom_editor_options" db:"custom_editor_options"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
}

type CustomEditorOptions struct {
	Theme       string `json:"theme,omitempty"`
	LineNumbers bool   `json:"line_numbers,omitempty"`
	WordWrap    bool   `json:"word_wrap,omitempty"`
	Folding     bool   `json:"folding,omitempty"`
	Minimap     bool   `json:"minimap,omitempty"`
	FontFamily  string `json:"font_family,omitempty"`
}

func (pc *CustomEditorOptions) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return nil
	}
}
func (pc CustomEditorOptions) Value() (driver.Value, error) {
	result, err := json.Marshal(pc)
	return string(result), err
}

//TableName - returns table name in database
func (s Snippet) TableName() string {
	return "snippets"
}

func (s *Snippet) Load(snippet Snippet) {
	s.UserID = snippet.UserID
	s.Favorite = snippet.Favorite
	s.AccessLevel = snippet.AccessLevel
	s.Title = snippet.Title
	s.Content = snippet.Content
	s.Language = snippet.Language
	s.CustomEditorOptions = snippet.CustomEditorOptions
	s.UpdatedAt = snippet.UpdatedAt
}
