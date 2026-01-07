-- name: CreateCharacter :one
INSERT INTO characters (
    id, name, description, level, experience, 
    armor_class, hit_points, strength, dexterity, 
    constitution, intelligence, wisdom, charisma
) VALUES (
    $1, $2, $3, $4, $5, 
    $6, $7, $8, $9, 
    $10, $11, $12, $13
) RETURNING *;

-- name: GetCharacter :one
SELECT 
    c.*,
    cl.name AS class_name,
    cc.class_level
FROM characters c
JOIN character_classes cc ON c.id = cc.character_id 
JOIN classes cl ON cc.class_id = cl.id;             

-- name: ListCharacters :many
SELECT * FROM characters 
ORDER BY name;

-- name: UpdateCharacterStats :exec
UPDATE characters
SET 
    level = $2,
    experience = $3,
    hit_points = $4,
    strength = $5,
    dexterity = $6,
    constitution = $7,
    intelligence = $8,
    wisdom = $9,
    charisma = $10
WHERE id = $1;

-- name: DeleteCharacter :exec
DELETE FROM characters WHERE id = $1;