package postgres

import (
	"context"
	"database/sql"
)

type PGArticle struct {
	DB *sql.DB
}

func (pg *PGArticle) Get(ctx context.Context, articleID *string) error {
	rows, err := pg.DB.QueryContext(ctx, "SELECT * FROM artigos WHERE cidade_id = $1", *articleID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(); err != nil {
			return err
		}
	}

	return nil
}
