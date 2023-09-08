package postgres

import (
	"context"
	"database/sql"

	"nossobr/domain/article"
)

type PGArticle struct {
	DB *sql.DB
}

func (pg *PGArticle) Get(ctx context.Context, uf, slug *string) (*article.Article, error) {
	res := new(article.Article)

	return res, pg.DB.QueryRowContext(ctx, `
		SELECT TA.id, TA.content, TA.city_id, TA.created_at, TA.updated_at
		FROM t_article TA
		JOIN t_cities TC ON TC.id = TA.city_id
		JOIN t_states TE ON TE.id = TC.state_id
		WHERE TC.slug = $1 AND TE.uf = $2`, slug, uf,
	).Scan(&res.ID, &res.Content, &res.CityID, &res.CreatedAt, &res.UpdatedAt)
}
