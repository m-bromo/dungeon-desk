package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/m-bromo/dungeon-desk/character-service/config"
	repository "github.com/m-bromo/dungeon-desk/character-service/internal/Repository"
	"github.com/m-bromo/dungeon-desk/character-service/internal/database"
	"github.com/m-bromo/dungeon-desk/character-service/internal/database/sqlc"
	"github.com/m-bromo/dungeon-desk/character-service/internal/service"
	"github.com/m-bromo/dungeon-desk/character-service/internal/web/handler"
	"github.com/m-bromo/dungeon-desk/character-service/internal/web/routes"
	"github.com/m-bromo/dungeon-desk/character-service/logger"
	"github.com/m-bromo/dungeon-desk/character-service/pkg/injector"
	"go.uber.org/dig"
)

func main() {
	config.SetupConfig()
	logger.SetupLog(config.Env.Environment)

	r := gin.Default()
	c := dig.New()

	injector.Provide(c, database.NewPostgresConnection)
	injector.Provide(c, sqlc.NewRepository)
	injector.Provide(c, repository.NewCharacterRepository)
	injector.Provide(c, service.NewCharacterService)
	injector.Provide(c, handler.NewCharacterHandler)

	routes.SetupRoutes(r, c)

	slog.Info("Starting service", "port", config.Env.Api.Host)
	r.Run(fmt.Sprintf("%s:%s", config.Env.Api.Host, config.Env.Api.Port))
}
