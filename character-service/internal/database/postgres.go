package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/m-bromo/dungeon-desk/character-service/config"
)

func NewPostgresConnection() (*sql.DB, error) {
	dns := fmt.Sprintf("name=%s user=%s password=%s host=%s port=%s sslmode=disabled",
		config.Env.PostgresDB.Name,
		config.Env.PostgresDB.User,
		config.Env.PostgresDB.Password,
		config.Env.PostgresDB.Host,
		config.Env.PostgresDB.Port,
	)

	conn, err := sql.Open("postgres", dns)
	if err != nil {
		slog.Error("Failed to connect to postgres database", "error", err)
		return nil, err
	}

	slog.Info("Connected to postgres database")

	return conn, nil
}
