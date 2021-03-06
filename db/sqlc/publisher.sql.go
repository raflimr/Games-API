// Code generated by sqlc. DO NOT EDIT.
// source: publisher.sql

package db

import (
	"context"
)

const createPublisher = `-- name: CreatePublisher :exec
INSERT INTO publisher(
    logopub, nama, deskripsi, website 
)VALUES(
    ?,?,?,?
)
`

type CreatePublisherParams struct {
	Logopub   string `json:"logopub"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Website   string `json:"website"`
}

func (q *Queries) CreatePublisher(ctx context.Context, arg CreatePublisherParams) error {
	_, err := q.db.ExecContext(ctx, createPublisher,
		arg.Logopub,
		arg.Nama,
		arg.Deskripsi,
		arg.Website,
	)
	return err
}

const deletePublisher = `-- name: DeletePublisher :exec
DELETE FROM publisher WHERE id =?
`

func (q *Queries) DeletePublisher(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePublisher, id)
	return err
}

const listPublisher = `-- name: ListPublisher :many
SELECT id, logopub, nama, deskripsi, website FROM publisher
`

func (q *Queries) ListPublisher(ctx context.Context) ([]Publisher, error) {
	rows, err := q.db.QueryContext(ctx, listPublisher)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Publisher
	for rows.Next() {
		var i Publisher
		if err := rows.Scan(
			&i.ID,
			&i.Logopub,
			&i.Nama,
			&i.Deskripsi,
			&i.Website,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePublisher = `-- name: UpdatePublisher :exec
UPDATE publisher SET logopub = ?, nama =?, deskripsi =?, website =? WHERE id =?
`

type UpdatePublisherParams struct {
	Logopub   string `json:"logopub"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Website   string `json:"website"`
	ID        int32  `json:"id"`
}

func (q *Queries) UpdatePublisher(ctx context.Context, arg UpdatePublisherParams) error {
	_, err := q.db.ExecContext(ctx, updatePublisher,
		arg.Logopub,
		arg.Nama,
		arg.Deskripsi,
		arg.Website,
		arg.ID,
	)
	return err
}
