-- name: CreateFlaw :one
INSERT INTO flaws (id, character_id, description) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFlawsByCharacterID :many
SELECT * FROM flaws 
WHERE character_id = $1
ORDER BY description;

-- name: DeleteFlaw :exec
DELETE FROM flaws 
WHERE id = $1;