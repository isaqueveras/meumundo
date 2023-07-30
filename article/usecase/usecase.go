package usecase

import (
	"context"
	"time"

	"nossobr/domain"
)

type articleUsecase struct {
	articleRepo    domain.IArticle
	contextTimeout time.Duration
}

func NewArticleUsecase(a domain.IArticle, timeout time.Duration) domain.ArticleUsecase {
	return &articleUsecase{articleRepo: a, contextTimeout: timeout}
}

func (a *articleUsecase) Get(ctx context.Context, uf, slug *string) (article *domain.Article, err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	if article, err = a.articleRepo.Get(ctx, uf, slug); err != nil {
		return nil, err
	}

	return
}
