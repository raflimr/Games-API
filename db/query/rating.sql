-- name: CreateRating :exec
INSERT INTO rating(
    rating, game_id, user_id
)VALUES(
    ?,?,?
);

-- name: UpdateRating :exec
UPDATE rating SET rating = ?, game_id =?, user_id =? WHERE id =?;