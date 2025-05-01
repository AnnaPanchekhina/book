package models

import "time"
import "errors"

type Book struct {
    ID        int64     `json:"id"`
    Title     string    `json:"title"`
    Author    string    `json:"author"`
    Year      int       `json:"year,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (b Book) Validate() error {
	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.Author == "" {
		return errors.New("author is required")
	}
	return nil
}