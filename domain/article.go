package domain

import "time"

// Article is representing the Article data struct
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    Author    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface{}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface{}
