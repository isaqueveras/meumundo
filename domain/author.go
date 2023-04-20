package domain

// IAuthor represent the author's repository contract
type IAuthor interface{}

// Author representing the Author data struct
type Author struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
