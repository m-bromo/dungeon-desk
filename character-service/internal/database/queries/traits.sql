-- name: CreateTrait :one
INSERT INTO traits (id, character_id, description) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTraitsByCharacterID :many
SELECT * FROM traits 
WHERE character_id = $1
ORDER BY description;

-- name: DeleteTrait :exec
DELETE FROM traits 
WHERE id = $1;