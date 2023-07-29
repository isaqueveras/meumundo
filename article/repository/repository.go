package repository

import (
	"context"
	"database/sql"
	"nossobr/article/repository/postgres"
	"nossobr/domain"
)

type repository struct {
	pg *postgres.PGArticle
}

func NewRepo(conn *sql.DB) domain.IArticle {
	return &repository{pg: &postgres.PGArticle{DB: conn}}
}

func (r *repository) Get(ctx context.Context, articleID *string) error {
	return r.pg.Get(ctx, articleID)
}
