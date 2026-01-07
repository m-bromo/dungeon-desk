package dto

type CreateCharacterClassPayload struct {
	ClassID int `json:"classId" binding:"required"`
	Level   int `json:"level" binding:"required,min=1"`
}
