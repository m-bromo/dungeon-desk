-- name: ListAllClasses :many
SELECT * FROM classes ORDER BY name;

-- name: AddClassToCharacter :one
INSERT INTO character_classes (character_id, class_id, class_level) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetClassesByCharacterID :many
SELECT c.id, c.name, cc.class_level
FROM classes c
JOIN character_classes cc ON c.id = cc.class_id
WHERE cc.character_id = $1;