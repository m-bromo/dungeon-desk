package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain/entities"
)

type ClassRepository interface {
	GetCharacterClasses(ctx context.Context, characterID uuid.UUID) ([]*entities.Class, error)
}
