package domain

import (
	"context"
	"time"
)

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Get(ctx context.Context, id *string) error
}

// IArticle represent the article's repository contract
type IArticle interface {
	Get(ctx context.Context, id *string) error
}

// Article is representing the Article data struct
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
