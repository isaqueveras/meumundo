package article

import (
	"context"
	"database/sql"

	"nossobr/domain/article"
	"nossobr/repository/article/postgres"
)

type repository struct {
	pg *postgres.PGArticle
}

func New(conn *sql.DB) article.IArticle {
	return &repository{pg: &postgres.PGArticle{DB: conn}}
}

func (r *repository) Get(ctx context.Context, uf, slug *string) (*article.Article, error) {
	return r.pg.Get(ctx, uf, slug)
}
