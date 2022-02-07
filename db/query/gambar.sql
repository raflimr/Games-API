-- name: GetGambar :one
SELECT url FROM gambar WHERE id =?;

-- name: CreateGambar :exec
INSERT INTO gambar(
    url, game_id
)VALUES(
    ?,?
);