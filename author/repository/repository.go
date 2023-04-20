package repository

import (
	"nossobr/author/repository/postgres"
	"nossobr/database"
	"nossobr/domain"
)

type repo struct {
	pg *postgres.PGAuthor
}

// New will create an implementation of author.Repository
func New(tx *database.DBTransaction) domain.IAuthor {
	return &repo{pg: &postgres.PGAuthor{DB: tx}}
}
