-- name: CreateReaksiReview :exec
INSERT INTO reaksi_review(
    tipe_reaksi,review_id, user_id 
)VALUES(
    ?,?,?
);

-- name: UpdateReaksiReview :exec
UPDATE reaksi_review SET tipe_reaksi =?, review_id =?, user_id =? WHERE id =?;

-- name: GetReaksiReview :many
SELECT * FROM reaksi_review;

