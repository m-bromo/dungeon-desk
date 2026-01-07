package mapper

import (
	"database/sql"

	db "github.com/m-bromo/dungeon-desk/character-service/internal/database/sqlc"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain"
)

func ToAddClassToCharacterParams(d *domain.CharacterClass) db.AddClassToCharacterParams {
	return db.AddClassToCharacterParams{
		CharacterID: d.CharacterID,
		ClassID:     int32(d.ClassID),
		ClassLevel: sql.NullInt32{
			Int32: int32(d.ClassLevel),
			Valid: true,
		},
	}
}
