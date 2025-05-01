package models

import "testing"

func TestBookValidation(t *testing.T) {
	tests := []struct {
		name    string
		book    Book
		wantErr bool
	}{
		{
			name:    "Valid Book",
			book:    Book{Title: "Good Book", Author: "Author"},
			wantErr: false,
		},
		{
			name:    "Empty Title",
			book:    Book{Title: "", Author: "Author"},
			wantErr: true,
		},
		{
			name:    "Empty Author",
			book:    Book{Title: "Title", Author: ""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.book.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}