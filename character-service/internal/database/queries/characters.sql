-- name: CreateCharacter :one
INSERT INTO characters (
    id, 
    name, 
    description, 
    alignment, 
    total_level, 
    experience, 
    armor_class, 
    hit_points, 
    current_hit_points, 
    strength, 
    dexterity, 
    constitution, 
    intelligence, 
    wisdom, 
    charisma, 
    traits, 
    flaws
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
) RETURNING *;         

-- name: ListCharacters :many
SELECT * FROM characters 
ORDER BY name;

-- name: DeleteCharacter :exec
DELETE FROM characters WHERE id = $1;