package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

//Snippet - ...
type Snippet struct {
	ID                  int                 `json:"id"`
	UserID              int                 `json:"user_id"`
	Favorite            bool                `json:"favorite"`
	AccessLevel         int                 `json:"access_level"`
	Title               string              `json:"title"`
	Content             string              `json:"content"`
	Tags                Tags                `json:"tags" db:"tags"`
	Language            string              `json:"language"`
	CustomEditorOptions CustomEditorOptions `json:"custom_editor_options" db:"custom_editor_options"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
}

type Tags []string

type CustomEditorOptions struct {
	Theme       string `json:"theme,omitempty"`
	LineNumbers bool   `json:"line_numbers,omitempty"`
	WordWrap    bool   `json:"word_wrap,omitempty"`
	Folding     bool   `json:"folding,omitempty"`
	Minimap     bool   `json:"minimap,omitempty"`
	FontFamily  string `json:"font_family,omitempty"`
}

func (t *Tags) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &t)
	case string:
		return json.Unmarshal([]byte(v), &t)
	default:
		return nil
	}
}

func (t Tags) Value() (driver.Value, error) {
	result, err := json.Marshal(t)
	return string(result), err
}

func (pc *CustomEditorOptions) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &pc)
	case string:
		return json.Unmarshal([]byte(v), &pc)
	default:
		return nil
	}
}
func (pc CustomEditorOptions) Value() (driver.Value, error) {
	result, err := json.Marshal(pc)
	return string(result), err
}

//TableName - returns table name from database
func (s Snippet) TableName() string {
	return "snippets"
}

func (s Snippet) GetOwnerID() int {
	return s.UserID
}

func (s Snippet) IsPublic() bool {
	if s.AccessLevel == 0 {
		return false
	} else {
		return true
	}
}
