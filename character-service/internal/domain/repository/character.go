package repository

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/m-bromo/dungeon-desk/character-service/internal/domain/entities"
)

type CharacterRepository interface {
	CreateCharacter(ctx context.Context, character *domain.Character) error
	AddClassToCharacter(ctx context.Context, characterID uuid.UUID, classID int) error
	GetCharacter(ctx context.Context, characterID uuid.UUID) (*domain.Character, error)
}
