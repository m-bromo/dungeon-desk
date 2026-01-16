package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/m-bromo/dungeon-desk/character-service/internal/mapper"
	"github.com/m-bromo/dungeon-desk/character-service/internal/service"
	"github.com/m-bromo/dungeon-desk/character-service/internal/web/models"
)

type CharacterHandler struct {
	characterService service.CharacterService
}

func NewCharacterHandler(characterService service.CharacterService) *CharacterHandler {
	return &CharacterHandler{
		characterService: characterService,
	}
}

func (h *CharacterHandler) CreateCharacter(c *gin.Context) {
	var payload models.CreateCharacterPayload
	if err := c.Bind(&payload); err != nil {
		c.Error(err)
		return
	}

	input := mapper.ToDomainInput(&payload)

	if err := h.characterService.CreateCharacter(
		c.Request.Context(),
		input,
	); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h *CharacterHandler) GetCharacter(c *gin.Context) {
	characterID, _ := c.Params.Get("id")

	character, err := h.characterService.GetCharacter(c.Request.Context(), uuid.MustParse(characterID))
	if err != nil {
		c.Error(err)
		return
	}

	response := mapper.ToCharacterResponse(character)

	c.JSON(http.StatusOK, response)

}
