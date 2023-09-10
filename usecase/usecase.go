package usecase

import (
	"context"
	"nossobr/domain"
	"time"
)

type usecase struct {
	repo    domain.IBr
	timeout time.Duration
}

func NewUsecase(a domain.IBr, timeout time.Duration) domain.Usecase {
	return &usecase{repo: a, timeout: timeout}
}

func (a *usecase) GetArticle(ctx context.Context, uf, slug *string) (article *domain.Article, err error) {
	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	if article, err = a.repo.GetArticle(ctx, uf, slug); err != nil {
		return nil, err
	}

	return
}

func (a *usecase) GetChildren(ctx context.Context, uf, slug *string) (*[]*domain.Children, error) {
	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	children, err := a.repo.GetChildren(ctx, uf, slug)
	if err != nil {
		return nil, err
	}

	return children, nil
}
