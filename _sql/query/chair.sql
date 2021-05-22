-- name: GetChair :one
SELECT * FROM chair WHERE id = ?;

-- name: CreateChair :exec
INSERT INTO chair(id, name, description, thumbnail, price, height, width, depth, color, features, kind, popularity, stock) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: ListChairWithCondtion :many
SELECT * FROM chair
WHERE 
    (price > sqlc.arg(min_price) OR sqlc.arg(min_price) IS NULL)
AND (price <= sqlc.arg(max_price) OR sqlc.arg(max_price) IS NULL)
AND (height > sqlc.arg(min_height) OR sqlc.arg(min_height) IS NULL)
AND (height <= sqlc.arg(max_height) OR sqlc.arg(max_height) IS NULL)
AND (width > sqlc.arg(min_width) OR sqlc.arg(min_width) IS NULL)
AND (width <= sqlc.arg(max_width) OR sqlc.arg(max_width) IS NULL)
AND (width > sqlc.arg(min_depth) OR sqlc.arg(min_depth) IS NULL)
AND (width <= sqlc.arg(max_depth) OR sqlc.arg(max_depth) IS NULL)
AND (kind = sqlc.arg(kind) OR sqlc.arg(kind) IS NULL)
AND (color = sqlc.arg(color) OR sqlc.arg(color) IS NULL)
AND (features = sqlc.arg(features) OR sqlc.arg(features) IS NULL)
ORDER BY popularity DESC, id ASC;

-- name: GetExistChair :one
SELECT * FROM chair WHERE id = ? AND stock > 0;

-- name: UpdateChairForPurchase :exec
UPDATE chair SET stock = stock - 1 WHERE id = ?;