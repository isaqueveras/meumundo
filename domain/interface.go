package domain

import "context"

// Usecase represent the article's usecases
type Usecase interface {
	GetArticle(ctx context.Context, uf, slug *string) (*Article, error)
	GetChildren(ctx context.Context, uf, slug *string) (*[]*Children, error)
}

// IBr represent the repository contract
type IBr interface {
	GetArticle(ctx context.Context, uf, slug *string) (*Article, error)
	GetChildren(ctx context.Context, uf, slug *string) (*[]*Children, error)
}
