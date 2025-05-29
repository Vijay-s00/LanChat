-- name: GetMessages :many
SELECT * FROM Messages;


-- name: InsertMessage :one
INSERT INTO Messages (name_from, name_to, message, created_at)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: InsertBulkMessages :copyfrom
INSERT INTO Messages (name_from, name_to, message, created_at) 
VALUES ($1, $2, $3, $4);
