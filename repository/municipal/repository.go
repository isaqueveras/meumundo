package municipal

import (
	"database/sql"
	"meumundo/domain/municipal"
)

type repository struct {
	pg *sql.DB
}

// New create a new repository
func New(conn *sql.DB) municipal.IRepo {
	return &repository{pg: conn}
}
