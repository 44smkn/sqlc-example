-- name: GetChair :one
SELECT * FROM chair WHERE id = ?;

-- name: CreateChair :exec
INSERT INTO chair(id, name, description, thumbnail, price, height, width, depth, color, features, kind, popularity, stock) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);

-- name: ListChairWithPriceCondtion :many
SELECT * FROM chair WHERE price >= ? AND price < ? AND stock > 0 ORDER BY popularity DESC, id ASC;

-- name: CountChairWithPriceCondition :one
SELECT COUNT(*) FROM chair WHERE price >= ? AND price < ?;

-- name: GetExistChair :one
SELECT * FROM chair WHERE id = ? AND stock > 0;

-- name: UpdateChairForPurchase :exec
UPDATE chair SET stock = stock - 1 WHERE id = ?;