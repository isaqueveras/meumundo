package repository

import (
	"context"
	"database/sql"
	"meumundo/domain"
)

type repo struct {
	pg *sql.DB
}

func New(conn *sql.DB) domain.IBr {
	return &repo{pg: conn}
}

func (r *repo) GetArticle(ctx context.Context, id *string) (*domain.Article, error) {
	res := new(domain.Article)

	query := `SELECT content, status, created_at, updated_at FROM t_article TA WHERE id = $1 AND status = 'Publish'`
	if err := r.pg.QueryRowContext(ctx, query, id).Scan(&res.Content, &res.Status, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return nil, err
	}

	rows, err := r.pg.QueryContext(ctx, `SELECT name, value FROM t_article_props WHERE article_id = $1 ORDER BY sortkey ASC`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		prop := &domain.Props{}
		if err = rows.Scan(&prop.Name, &prop.Value); err != nil {
			return nil, err
		}
		res.Props = append(res.Props, prop)
	}

	_ = rows.Close()

	return res, nil
}

func (r *repo) GetChildren(ctx context.Context, id *string) (*[]*domain.Children, error) {
	query := `
		SELECT TC2."name", TC2.url, TC2.short_desc
		FROM public.t_cities TC
		JOIN public.t_states TS ON TS.id = TC.state_id
		JOIN public.t_children TC2 ON TC2.city_id = TC.id 
		WHERE TC.slug = $1 AND TS.uf = $2`

	q, err := r.pg.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	resp := []*domain.Children{}
	for q.Next() {
		children := &domain.Children{}
		if err := q.Scan(&children.Name, &children.URL, &children.ShortDesc); err != nil {
			return nil, err
		}
		resp = append(resp, children)
	}

	return &resp, nil
}

func (r *repo) GetBorderTowns(ctx context.Context, id *string) (*[][3]*string, error) {
	query := `WITH towns AS (
			SELECT unnest(TC.border_towns_id) AS cities
			FROM public.t_cities TC 
			JOIN public.t_states ts ON TS.id = TC.state_id 
			WHERE TC.slug = $1 AND TS.uf = $2
		) SELECT TC.slug, TC.city, TS.uf
		FROM public.t_cities TC 
		JOIN public.t_states TS ON TS.id = TC.state_id 
		JOIN towns T ON TC.id = T.cities`

	q, err := r.pg.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	resp := [][3]*string{}
	for q.Next() {
		var slug, city, uf *string
		if err := q.Scan(&slug, &city, &uf); err != nil {
			return nil, err
		}
		resp = append(resp, [3]*string{slug, city, uf})
	}

	return &resp, nil
}
