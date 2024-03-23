-- name: GetBoardRoomById :one
SELECT * FROM boardRooms
WHERE id = $1 LIMIT 1;

-- name: GetAllBoardRooms :many
SELECT * FROM boardRooms
ORDER BY createdAt;

-- name: CreateBoardRoom :one
INSERT INTO boardRooms (
  id, name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateBoardRoomById :one
UPDATE boardRooms
  set name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteBoardRoomById :exec
DELETE FROM boardRooms
WHERE id = $1;