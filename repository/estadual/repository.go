package estadual

import (
	"database/sql"
	"meumundo/domain/estadual"
)

type repository struct {
	pg *sql.DB
}

// New create a new repository
func New(conn *sql.DB) estadual.IRepo {
	return &repository{pg: conn}
}
