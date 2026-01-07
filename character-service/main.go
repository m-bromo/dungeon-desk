package main

import (
	"log"
	"log/slog"

	"github.com/m-bromo/dungeon-desk/character-service/config"
	"github.com/m-bromo/dungeon-desk/character-service/internal/database"
	"github.com/m-bromo/dungeon-desk/character-service/logger"
)

func main() {
	logger.SetupLog(config.Env.Environment)

	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()

	slog.Info("Starting service", "port", config.Env.Api.Host)
}
