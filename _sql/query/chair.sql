-- name: GetChair :one
SELECT * FROM chair WHERE id = ?;

-- name: CreateChair :exec
INSERT INTO chair(id, name, description, thumbnail, price, height, width, depth, color, features, kind, popularity, stock) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: ListChairWithCondtion :many
SELECT * FROM chair
WHERE 
    (price > sqlc.arg(min_price1) OR ISNULL(sqlc.arg(min_price2)))
AND (price <= sqlc.arg(max_price1) OR ISNULL(sqlc.arg(max_price2)))
AND (height > sqlc.arg(min_height1) OR ISNULL(sqlc.arg(min_height2)))
AND (height <= sqlc.arg(max_height1) OR ISNULL(sqlc.arg(max_height2)))
AND (width > sqlc.arg(min_width1) OR ISNULL(sqlc.arg(min_width2)))
AND (width <= sqlc.arg(max_width1) OR ISNULL(sqlc.arg(max_width2)))
AND (width > sqlc.arg(min_depth1) OR ISNULL(sqlc.arg(min_depth2)))
AND (width <= sqlc.arg(max_depth1) OR ISNULL(sqlc.arg(max_depth2)))
AND (kind = sqlc.arg(kind1) OR ISNULL(sqlc.arg(kind2)))
AND (color = sqlc.arg(color1) OR ISNULL(sqlc.arg(color2)))
AND (features = sqlc.arg(features1) OR ISNULL(sqlc.arg(features2)))
ORDER BY popularity DESC, id ASC;

-- name: GetExistChair :one
SELECT * FROM chair WHERE id = ? AND stock > 0;

-- name: UpdateChairForPurchase :exec
UPDATE chair SET stock = stock - 1 WHERE id = ?;