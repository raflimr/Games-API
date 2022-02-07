-- name: CreateKategori :exec
INSERT INTO kategori(
    nama, deskripsi
)VALUES(
    ?,?
);

-- name: UpdateKategori :exec
UPDATE kategori SET nama = ?, deskripsi =? WHERE id =?;

-- name: DeleteKategori :exec
DELETE FROM kategori WHERE id =?;

-- name: GetKategori :many
SELECT * FROM kategori;