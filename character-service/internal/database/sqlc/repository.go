package sqlc

import "database/sql"

func NewRepository(db *sql.DB) Querier {
	return New(db)
}
