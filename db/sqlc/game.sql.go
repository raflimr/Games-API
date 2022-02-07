// Code generated by sqlc. DO NOT EDIT.
// source: game.sql

package db

import (
	"context"
)

const createGames = `-- name: CreateGames :exec
INSERT INTO game(
    judul, deskripsi, penerbit, platform, website, status_game, tanggal_rilis, developer_id, publisher_id
)VALUES(
    ?,?,?,?,?,?,?,?,?
)
`

type CreateGamesParams struct {
	Judul        string `json:"judul"`
	Deskripsi    string `json:"deskripsi"`
	Penerbit     string `json:"penerbit"`
	Platform     string `json:"platform"`
	Website      string `json:"website"`
	StatusGame   string `json:"status_game"`
	TanggalRilis string `json:"tanggal_rilis"`
	DeveloperID  int32  `json:"developer_id"`
	PublisherID  int32  `json:"publisher_id"`
}

func (q *Queries) CreateGames(ctx context.Context, arg CreateGamesParams) error {
	_, err := q.db.ExecContext(ctx, createGames,
		arg.Judul,
		arg.Deskripsi,
		arg.Penerbit,
		arg.Platform,
		arg.Website,
		arg.StatusGame,
		arg.TanggalRilis,
		arg.DeveloperID,
		arg.PublisherID,
	)
	return err
}

const deleteGames = `-- name: DeleteGames :exec
DELETE FROM game WHERE id = ?
`

func (q *Queries) DeleteGames(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteGames, id)
	return err
}

const listGames = `-- name: ListGames :many
SELECT
    g.id,
    g.judul,
    g.deskripsi,
    g.penerbit,
    g.platform,
    g.website,
    g.status_game,
    g.tanggal_rilis,
 COALESCE(GROUP_CONCAT(DISTINCT k.nama), "") kategori,
    d.pendiri developer,
    p.nama publisher,
    COALESCE(AVG(r.rating), 0) rating,
    COALESCE(GROUP_CONCAT(DISTINCT gb.url),"") gambar 
FROM
    game g
LEFT JOIN
    rating r ON g.id = r.game_id
JOIN
    developer d ON g.developer_id = d.id
JOIN
    publisher p ON g.publisher_id = p.id
LEFT JOIN
    kategori_game kg ON g.id = kg.game_id
LEFT JOIN
    kategori k ON kg.kategori_id = k.id
LEFT JOIN
  gambar gb ON g.id = gb.game_id 
GROUP BY
    g.id
`

type ListGamesRow struct {
	ID           int32   `json:"id"`
	Judul        string  `json:"judul"`
	Deskripsi    string  `json:"deskripsi"`
	Penerbit     string  `json:"penerbit"`
	Platform     string  `json:"platform"`
	Website      string  `json:"website"`
	StatusGame   string  `json:"status_game"`
	TanggalRilis string  `json:"tanggal_rilis"`
	Kategori     string  `json:"kategori"`
	Developer    string  `json:"developer"`
	Publisher    string  `json:"publisher"`
	Rating       float32 `json:"rating"`
	Gambar       string  `json:"gambar"`
}

func (q *Queries) ListGames(ctx context.Context) ([]ListGamesRow, error) {
	rows, err := q.db.QueryContext(ctx, listGames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListGamesRow
	for rows.Next() {
		var i ListGamesRow
		if err := rows.Scan(
			&i.ID,
			&i.Judul,
			&i.Deskripsi,
			&i.Penerbit,
			&i.Platform,
			&i.Website,
			&i.StatusGame,
			&i.TanggalRilis,
			&i.Kategori,
			&i.Developer,
			&i.Publisher,
			&i.Rating,
			&i.Gambar,
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

const updateGames = `-- name: UpdateGames :exec
UPDATE game SET judul = ?, deskripsi = ?, penerbit = ?, platform = ?, website = ?, status_game = ?, tanggal_rilis = ? WHERE id = ?
`

type UpdateGamesParams struct {
	Judul        string `json:"judul"`
	Deskripsi    string `json:"deskripsi"`
	Penerbit     string `json:"penerbit"`
	Platform     string `json:"platform"`
	Website      string `json:"website"`
	StatusGame   string `json:"status_game"`
	TanggalRilis string `json:"tanggal_rilis"`
	ID           int32  `json:"id"`
}

func (q *Queries) UpdateGames(ctx context.Context, arg UpdateGamesParams) error {
	_, err := q.db.ExecContext(ctx, updateGames,
		arg.Judul,
		arg.Deskripsi,
		arg.Penerbit,
		arg.Platform,
		arg.Website,
		arg.StatusGame,
		arg.TanggalRilis,
		arg.ID,
	)
	return err
}