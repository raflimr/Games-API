-- name: CreateDeveloper :exec
INSERT INTO developer(
    logodev, kantor_pusat, pendiri, tahun_pendirian 
)VALUES(    
    ?,?,?,?
);

-- name: ListDeveloper :many
SELECT * FROM developer;

-- name: UpdateDeveloper :exec
UPDATE developer SET logodev = ?, kantor_pusat = ?, pendiri =?, tahun_pendirian =? WHERE id =?;

-- name: DeleteDeveloper :exec
DELETE FROM developer WHERE id =?;
