package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/m-bromo/dungeon-desk/character-service/config"
	"github.com/m-bromo/dungeon-desk/character-service/internal/database"
	"github.com/m-bromo/dungeon-desk/character-service/internal/database/sqlc"
	"github.com/m-bromo/dungeon-desk/character-service/logger"
	"github.com/m-bromo/dungeon-desk/character-service/pkg/injector"
	"go.uber.org/dig"
)

func main() {
	config.SetupConfig()
	logger.SetupLog(config.Env.Environment)

	r := gin.Default()
	d := dig.New()

	injector.Provide(d, database.NewPostgresConnection)
	injector.Provide(d, sqlc.New)

	slog.Info("Starting service", "port", config.Env.Api.Host)
	r.Run(fmt.Sprintf("%s:%s", config.Env.Api.Host, config.Env.Api.Port))
}
