package usecase

import (
	"context"
	"time"

	"nossobr/domain/article"
)

type articleUsecase struct {
	articleRepo    article.IArticle
	contextTimeout time.Duration
}

func NewArticleUsecase(a article.IArticle, timeout time.Duration) article.ArticleUsecase {
	return &articleUsecase{articleRepo: a, contextTimeout: timeout}
}

func (a *articleUsecase) Get(ctx context.Context, uf, slug *string) (article *article.Article, err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	if article, err = a.articleRepo.Get(ctx, uf, slug); err != nil {
		return nil, err
	}

	return
}
