package repository

import (
	"nossobr/article/repository/postgres"
	"nossobr/database"
	"nossobr/domain"
)

type repo struct {
	pg *postgres.PGArticle
}

// New will create an object that represent the article.Repository interface
func New(tx *database.DBTransaction) domain.IArticle {
	return &repo{pg: &postgres.PGArticle{DB: tx}}
}
