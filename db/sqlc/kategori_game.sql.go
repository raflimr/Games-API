// Code generated by sqlc. DO NOT EDIT.
// source: kategori_game.sql

package db

import (
	"context"
)

const createGameKategori = `-- name: CreateGameKategori :exec
INSERT INTO kategori_game(
  kategori_id, game_id
)VALUES(
  ?,?
)
`

type CreateGameKategoriParams struct {
	KategoriID int32 `json:"kategori_id"`
	GameID     int32 `json:"game_id"`
}

func (q *Queries) CreateGameKategori(ctx context.Context, arg CreateGameKategoriParams) error {
	_, err := q.db.ExecContext(ctx, createGameKategori, arg.KategoriID, arg.GameID)
	return err
}

const getGamesByKategori = `-- name: GetGamesByKategori :many
SELECT
  g.id, g.judul, g.deskripsi, g.penerbit, g.platform, g.website, g.status_game, g.tanggal_rilis, g.developer_id, g.publisher_id 
FROM
  kategori_game kg
JOIN
  kategori k ON kg.kategori_id = k.id
JOIN
  game g ON kg.game_id = g.id
WHERE
  kg.kategori_id = ?
`

func (q *Queries) GetGamesByKategori(ctx context.Context, kategoriID int32) ([]Game, error) {
	rows, err := q.db.QueryContext(ctx, getGamesByKategori, kategoriID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Game
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.Judul,
			&i.Deskripsi,
			&i.Penerbit,
			&i.Platform,
			&i.Website,
			&i.StatusGame,
			&i.TanggalRilis,
			&i.DeveloperID,
			&i.PublisherID,
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
