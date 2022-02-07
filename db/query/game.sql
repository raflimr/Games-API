-- name: ListGames :many
SELECT
    g.id,
    g.judul,
    g.deskripsi,
    g.penerbit,
    g.platform,
    g.website,
    g.status_game,
    g.tanggal_rilis,
    CONCAT(k.nama) kategori,
    d.pendiri developer,
    p.nama publisher,
    AVG(r.rating) rating,
    CONCAT(gb.url) gambar
FROM
    game g
JOIN
    developer d ON g.developer_id = d.id
JOIN
    publisher p ON g.publisher_id = p.id
LEFT JOIN
    kategori_game kg ON g.id = kg.game_id
LEFT JOIN 
  kategori k ON kg.kategori_id = k.id
LEFT JOIN
    rating r ON g.id = r.game_id
LEFT JOIN
  gambar gb ON g.id = gb.game_id 
GROUP BY
  g.id;

-- name: CreateGames :exec
INSERT INTO game(
    judul, deskripsi, penerbit, platform, website, status_game, tanggal_rilis, developer_id, publisher_id
)VALUES(
    ?,?,?,?,?,?,?,?,?
);

-- name: UpdateGames :exec
UPDATE game SET judul = ?, deskripsi = ?, penerbit = ?, platform = ?, website = ?, status_game = ?, tanggal_rilis = ? WHERE id = ?;
    
-- name: DeleteGames :exec
DELETE FROM game WHERE id = ?;