-- name: CreatePublisher :exec
INSERT INTO publisher(
    logopub, nama, deskripsi, website 
)VALUES(
    ?,?,?,?
);

-- name: ListPublisher :many
SELECT * FROM publisher;

-- name: UpdatePublisher :exec
UPDATE publisher SET logopub = ?, nama =?, deskripsi =?, website =? WHERE id =?;

-- name: DeletePublisher :exec
DELETE FROM publisher WHERE id =?;
