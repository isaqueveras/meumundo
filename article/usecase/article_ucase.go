package usecase

import (
	"time"

	"nossobr/domain"
)

type articleUsecase struct {
	articleRepo    domain.ArticleRepository
	authorRepo     domain.AuthorRepository
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewArticleUsecase(a domain.ArticleRepository, ar domain.AuthorRepository, timeout time.Duration) domain.ArticleUsecase {
	return &articleUsecase{articleRepo: a, authorRepo: ar, contextTimeout: timeout}
}
