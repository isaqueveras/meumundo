package federal

import (
	"database/sql"
	"meumundo/domain/federal"
)

type repository struct {
	pg *sql.DB
}

// New create a new repository
func New(conn *sql.DB) federal.IRepo {
	return &repository{pg: conn}
}
