package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/m-bromo/dungeon-desk/character-service/config"
	"github.com/m-bromo/dungeon-desk/character-service/internal/database"
	"github.com/m-bromo/dungeon-desk/character-service/logger"
)

func main() {
	config.SetupConfig()
	logger.SetupLog(config.Env.Environment)

	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()

	r := gin.Default()

	slog.Info("Starting service", "port", config.Env.Api.Host)
	r.Run(fmt.Sprintf("%s:%s", config.Env.Api.Host, config.Env.Api.Port))
}
