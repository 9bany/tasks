// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: sites.sql

package db

import (
	"context"
	"encoding/json"
)

const createSite = `-- name: CreateSite :one
INSERT INTO sites (
  url, meta_data
) VALUES (
  $1, $2
)
RETURNING id, url, meta_data, created_at
`

type CreateSiteParams struct {
	Url      string          `json:"url"`
	MetaData json.RawMessage `json:"meta_data"`
}

func (q *Queries) CreateSite(ctx context.Context, arg CreateSiteParams) (Sites, error) {
	row := q.db.QueryRowContext(ctx, createSite, arg.Url, arg.MetaData)
	var i Sites
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.MetaData,
		&i.CreatedAt,
	)
	return i, err
}

const getSiteByURL = `-- name: GetSiteByURL :one
SELECT id, url, meta_data, created_at FROM sites
WHERE url = $1
`

func (q *Queries) GetSiteByURL(ctx context.Context, url string) (Sites, error) {
	row := q.db.QueryRowContext(ctx, getSiteByURL, url)
	var i Sites
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.MetaData,
		&i.CreatedAt,
	)
	return i, err
}
