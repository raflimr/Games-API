-- name: GetGamesByKategori :many
SELECT
  g.* 
FROM
  kategori_game kg
JOIN
  kategori k ON kg.kategori_id = k.id
JOIN
  game g ON kg.game_id = g.id
WHERE
  kg.kategori_id = ?;

-- name: CreateGameKategori :exec
INSERT INTO kategori_game(
  kategori_id, game_id
)VALUES(
  ?,?
);