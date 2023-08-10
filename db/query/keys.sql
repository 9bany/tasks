-- name: CreateKey :one
INSERT INTO keys (
  key
) VALUES (
  $1
)
RETURNING *;

-- name: GetKey :one
SELECT * FROM keys
WHERE key = $1;

-- name: GetRandomKey :one
SELECT * FROM keys
ORDER BY RANDOM()
LIMIT 1;

-- name: IncreaseKeyUsageCount :exec
UPDATE keys 
SET usage_count = usage_count + 1
WHERE id = $1;


