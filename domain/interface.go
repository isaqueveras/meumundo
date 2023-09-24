package domain

import "context"

// Usecase represent the article's usecases
type Usecase interface {
	GetArticle(ctx context.Context, id *string) (*Article, error)
	GetChildren(ctx context.Context, id *string) (*[]*Children, error)
	GetBorderTowns(ctx context.Context, id *string) (*[][3]*string, error)
}

// IBr represent the repository contract
type IBr interface {
	GetArticle(ctx context.Context, id *string) (*Article, error)
	GetChildren(ctx context.Context, id *string) (*[]*Children, error)
	GetBorderTowns(ctx context.Context, id *string) (*[][3]*string, error)
}
