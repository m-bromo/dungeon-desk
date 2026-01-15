package service

import (
	"context"
	"log/slog"

	repository "github.com/m-bromo/dungeon-desk/character-service/internal/Repository"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain"
)

type CharacterService interface {
	CreateCharacter(ctx context.Context, input *domain.CreateCharacterInput) error
}

type characterService struct {
	repository repository.CharacterRepository
}

func NewCharacterService(repository repository.CharacterRepository) CharacterService {
	return &characterService{
		repository: repository,
	}
}

func (s *characterService) CreateCharacter(ctx context.Context, input *domain.CreateCharacterInput) error {
	character := domain.NewCharacter(input)

	if err := s.repository.CreateCharacter(ctx, character); err != nil {
		slog.Error("Failed to create character", "error", err.Error())
		return err
	}

	if err := s.repository.AddClassToCharacter(ctx, character.ID, input.ClassID); err != nil {
		slog.Error("Failed to add class to character")
		return err
	}

	return nil
}
