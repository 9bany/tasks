-- name: CreateSite :one
INSERT INTO sites (
  url, meta_data
) VALUES (
  $1, $2
)
RETURNING *;


-- name: GetSiteByURL :one
SELECT * FROM sites
WHERE url = $1;
