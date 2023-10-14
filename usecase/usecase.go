package usecase

import (
	"context"
	"meumundo/domain"
	"time"
)

type usecase struct {
	repo    domain.IBr
	timeout time.Duration
}

func NewUsecase(a domain.IBr, timeout time.Duration) domain.Usecase {
	return &usecase{repo: a, timeout: timeout}
}

func (a *usecase) GetArticle(ctx context.Context, id *string) (article *domain.Article, err error) {
	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	if article, err = a.repo.GetArticle(ctx, id); err != nil {
		return nil, err
	}

	return
}

func (a *usecase) GetChildren(ctx context.Context, id *string) (*[]*domain.Children, error) {
	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	children, err := a.repo.GetChildren(ctx, id)
	if err != nil {
		return nil, err
	}

	return children, nil
}

func (a *usecase) GetBorderTowns(ctx context.Context, id *string) (*[][3]*string, error) {
	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	data, err := a.repo.GetBorderTowns(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
