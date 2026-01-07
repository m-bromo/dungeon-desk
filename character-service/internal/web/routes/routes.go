package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/m-bromo/dungeon-desk/character-service/internal/web/handler"
	"github.com/m-bromo/dungeon-desk/character-service/internal/web/middleware"
	"github.com/m-bromo/dungeon-desk/character-service/pkg/injector"
	"go.uber.org/dig"
)

func SetupRoutes(router *gin.Engine, c *dig.Container) {
	h := injector.Resolve[*handler.CharacterHandler](c)

	router.POST("/character/create", h.CreateCharacter, middleware.ErrorhandlerMiddleware)
}
