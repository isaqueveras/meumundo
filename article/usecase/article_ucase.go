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

func (a *articleUsecase) Get(ctx context.Context, articleID *string) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	if err := a.articleRepo.Get(ctx, articleID); err != nil {
		return err
	}

	return nil
}
