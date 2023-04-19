package domain

// Author representing the Author data struct
type Author struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AuthorRepository represent the author's repository contract
type AuthorRepository interface{}
