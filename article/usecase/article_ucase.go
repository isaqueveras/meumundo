package usecase

import (
	"time"

	"nossobr/domain"
)

type articleUsecase struct {
	articleRepo    domain.IArticle
	authorRepo     domain.IAuthor
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewArticleUsecase(a domain.IArticle, ar domain.IAuthor, timeout time.Duration) domain.ArticleUsecase {
	return &articleUsecase{articleRepo: a, authorRepo: ar, contextTimeout: timeout}
}
