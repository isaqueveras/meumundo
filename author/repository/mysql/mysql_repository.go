package mysql

import (
	"context"
	"database/sql"

	"nossobr/domain"
)

type mysqlAuthorRepo struct {
	DB *sql.DB
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewMysqlAuthorRepository(db *sql.DB) domain.AuthorRepository {
	return &mysqlAuthorRepo{DB: db}
}

func (m *mysqlAuthorRepo) getOne(ctx context.Context, query string, args ...interface{}) (res domain.Author, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.Author{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.Author{}

	err = row.Scan(&res.ID, &res.Name, &res.CreatedAt, &res.UpdatedAt)
	return
}

func (m *mysqlAuthorRepo) GetByID(ctx context.Context, id int64) (domain.Author, error) {
	return m.getOne(ctx, `SELECT id, name, created_at, updated_at FROM author WHERE id=?`, id)
}
