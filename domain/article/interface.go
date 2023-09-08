package article

import "context"

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Get(ctx context.Context, uf, slug *string) (*Article, error)
}

// IArticle represent the article's repository contract
type IArticle interface {
	Get(ctx context.Context, uf, slug *string) (*Article, error)
}
