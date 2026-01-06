package main

import (
	"log/slog"

	"github.com/m-bromo/dungeon-desk/character-service/config"
	"github.com/m-bromo/dungeon-desk/character-service/logger"
)

func main() {
	logger.SetupLog(config.Env.Environment)
	slog.Info("Starting service", "port", config.Env.Api.Host)
}
