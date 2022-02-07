-- name: CreateReview :exec
INSERT INTO review(
    konten_review, tanggal_review, tipe, game_id, user_id
)VALUES(
    ?,now(),?,?,?
);

-- name: UpdateReview :exec
UPDATE review SET konten_review = ?, tanggal_review = ?, tipe = ?, game_id = ?, user_id = ? WHERE id =?;

-- name: DeleteReview :exec
DELETE FROM review WHERE id = ?;

-- name: GetReview :many
SELECT
    r.id,
    r.konten_review,
    r.tanggal_review,
    r.tipe,
    r.game_id,
    r.user_id
FROM
    review r;