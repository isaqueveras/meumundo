package repository

import (
	"context"
	"database/sql"
	"nossobr/domain"
)

type repository struct {
	pg *sql.DB
}

func New(conn *sql.DB) domain.IBr {
	return &repository{pg: conn}
}

func (r *repository) GetArticle(ctx context.Context, uf, slug *string) (*domain.Article, error) {
	res := new(domain.Article)

	query := `
		SELECT TA.id, TA.content, TA.city_id, TA.created_at, TA.updated_at
		FROM t_article TA
		JOIN t_cities TC ON TC.id = TA.city_id
		JOIN t_states TE ON TE.id = TC.state_id
		WHERE TC.slug = $1 AND TE.uf = $2`

	q := r.pg.QueryRowContext(ctx, query, slug, uf)
	if err := q.Scan(&res.ID, &res.Content, &res.CityID, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) GetChildren(ctx context.Context, uf, slug *string) (*[]*domain.Children, error) {
	return nil, nil
}
