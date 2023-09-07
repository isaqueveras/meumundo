package domain

import (
	"context"
	"time"
)

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Get(ctx context.Context, uf, slug *string) (*Article, error)
}

// IArticle represent the article's repository contract
type IArticle interface {
	Get(ctx context.Context, uf, slug *string) (*Article, error)
}

// Article is representing the Article data struct
type Article struct {
	ID        *string    `json:"id,omitempty"`
	CityID    *string    `json:"city_id,omitempty"`
	Content   *string    `json:"content,omitempty"`
	Children  []*People  `json:"children,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// People models the data of an illustrious son of a city
type People struct {
	ID          *string    `json:"id,omitempty"`
	URL         *string    `json:"url,omitempty"`
	Name        *string    `json:"name,omitempty"`
	ShortDesc   *string    `json:"short_desc,omitempty"`
	Biography   *string    `json:"biography,omitempty"`
	Professions []*string  `json:"professions,omitempty"`
	Parents     [2]*People `json:"parents,omitempty"`
	DateBirth   *time.Time `json:"date_birth,omitempty"`
	DateDeath   *time.Time `json:"date_death,omitempty"`
}
