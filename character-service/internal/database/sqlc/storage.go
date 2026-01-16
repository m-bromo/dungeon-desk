package sqlc

import "database/sql"

func NewStorage(db *sql.DB) Querier {
	return New(db)
}
