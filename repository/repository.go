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
		SELECT TA.id, TA.content, TA.city_id, TA.created_at, TA.updated_at, TA.status
		FROM t_article TA
		JOIN t_cities TC ON TC.id = TA.city_id
		JOIN t_states TE ON TE.id = TC.state_id
		WHERE TC.slug = $1 AND TE.uf = $2 AND TA.status = 'Publish'`

	q := r.pg.QueryRowContext(ctx, query, slug, uf)
	if err := q.Scan(&res.ID, &res.Content, &res.CityID, &res.CreatedAt, &res.UpdatedAt, &res.Status); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) GetChildren(ctx context.Context, uf, slug *string) (*[]*domain.Children, error) {
	query := `
		SELECT TC2."name", TC2.url, TC2.short_desc
		FROM public.t_cities TC
		JOIN public.t_states TS ON TS.id = TC.state_id
		JOIN public.t_children TC2 ON TC2.city_id = TC.id 
		WHERE TC.slug = $1 AND TS.uf = $2`

	q, err := r.pg.QueryContext(ctx, query, slug, uf)
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

func (r *repository) GetBorderTowns(ctx context.Context, uf, slug *string) (*[][3]*string, error) {
	query := `WITH towns AS (
			SELECT unnest(TC.border_towns) AS cities
			FROM public.t_cities TC 
			JOIN public.t_states ts ON TS.id = TC.state_id 
			WHERE TC.slug = $1 AND TS.uf = $2
		) SELECT TC.slug, TC.city, TS.uf
		FROM public.t_cities TC 
		JOIN public.t_states TS ON TS.id = TC.state_id 
		JOIN towns T ON TC.id = T.cities`

	q, err := r.pg.QueryContext(ctx, query, slug, uf)
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
