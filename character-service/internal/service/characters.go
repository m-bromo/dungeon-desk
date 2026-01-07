package service

import (
	"context"
	"log/slog"

	"github.com/m-bromo/dungeon-desk/character-service/internal/database/sqlc"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain"
	"github.com/m-bromo/dungeon-desk/character-service/internal/mapper"
)

type CharacterService interface {
	CreateCharacter(ctx context.Context, character *domain.Character) error
}

type characterService struct {
	repository sqlc.Querier
}

func NewCharacterService(repository sqlc.Querier) CharacterService {
	return &characterService{
		repository: repository,
	}
}

func (s *characterService) CreateCharacter(ctx context.Context, character *domain.Character) error {
	if character.Level > domain.MaxLevel {
		slog.Warn("Character at max level", "error", domain.ErrCharacterAtMaxLevel)
		return domain.ErrCharacterAtMaxLevel
	}

	char, err := s.repository.CreateCharacter(ctx, mapper.ToCreateCharacterParams(character))
	if err != nil {
		slog.Error("Failed to create character", "error", err.Error())
		return err
	}

	for _, trait := range character.Traits {
		_, err := s.repository.CreateTrait(ctx, mapper.ToCreateTraitParams(trait, char.ID))
		if err != nil {
			slog.Error("Failed to create trait", "error", err.Error())
			return err
		}
	}

	for _, flaw := range character.Flaws {
		_, err := s.repository.CreateFlaw(ctx, mapper.ToCreateFlawParams(flaw, char.ID))
		if err != nil {
			slog.Error("Failed to create flaw", "error", err.Error())
			return err
		}
	}

	for _, class := range character.Classes {
		class.CharacterID = char.ID
		_, err := s.repository.AddClassToCharacter(ctx, mapper.ToAddClassToCharacterParams(class))
		if err != nil {
			slog.Error("Failed to add class to character", "error", err.Error())
			return err
		}
	}

	return nil
}
