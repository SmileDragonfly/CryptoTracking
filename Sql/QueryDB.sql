-- name: InsertBUSDPrice :exec
INSERT INTO tblBUSDPrice ("price") VALUES ($1);
-- name: GetLastestBUSDPrice :one
SELECT "price" FROM tblBUSDPrice ORDER BY "time" DESC LIMIT 1;
-- name: Get1MinAgoBUSDPrice :one
SELECT "price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '61 seconds') ORDER BY "time" ASC LIMIT 1;
-- name: Get5MinAgoBUSDPrice :one
SELECT "price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '301 seconds') ORDER BY "time" ASC LIMIT 1;
-- name: Get10MinAgoBUSDPrice :one
SELECT "price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '601 seconds') ORDER BY "time" ASC LIMIT 1;
-- name: Get15MinAgoBUSDPrice :one
SELECT "price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '901 seconds') ORDER BY "time" ASC LIMIT 1;
-- name: Get30MinAgoBUSDPrice :one
SELECT "price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '1801 seconds') ORDER BY "time" ASC LIMIT 1;
-- name: Get60MinAgoBUSDPrice :one
SELECT "price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '3601 seconds') ORDER BY "time" ASC LIMIT 1;

-- name: Insert1MinBUSDPercent :exec
INSERT INTO tblBUSD1MinPercent ("symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4);
-- name: Insert5MinBUSDPercent :exec
INSERT INTO tblBUSD5MinPercent ("symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4);
-- name: Insert10MinBUSDPercent :exec
INSERT INTO tblBUSD10MinPercent ("symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4);
-- name: Insert15MinBUSDPercent :exec
INSERT INTO tblBUSD15MinPercent ("symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4);
-- name: Insert30MinBUSDPercent :exec
INSERT INTO tblBUSD30MinPercent ("symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4);
-- name: Insert60MinBUSDPercent :exec
INSERT INTO tblBUSD60MinPercent ("symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4);

-- name: Delete1MinBUSDPercent :exec
DELETE FROM tblBUSD1MinPercent;
-- name: Delete5MinBUSDPercent :exec
DELETE FROM tblBUSD5MinPercent;
-- name: Delete10MinBUSDPercent :exec
DELETE FROM tblBUSD10MinPercent;
-- name: Delete15MinBUSDPercent :exec
DELETE FROM tblBUSD15MinPercent;
-- name: Delete30MinBUSDPercent :exec
DELETE FROM tblBUSD30MinPercent;
-- name: Delete60MinBUSDPercent :exec
DELETE FROM tblBUSD60MinPercent;

-- name: GetAll1MinPercent :many
SELECT * FROM tblBUSD1MinPercent;
-- name: GetAll5MinPercent :many
SELECT * FROM tblBUSD5MinPercent;
-- name: GetAll10MinPercent :many
SELECT * FROM tblBUSD10MinPercent;
-- name: GetAll15MinPercent :many
SELECT * FROM tblBUSD15MinPercent;
-- name: GetAll30MinPercent :many
SELECT * FROM tblBUSD30MinPercent;
-- name: GetAll60MinPercent :many
SELECT * FROM tblBUSD60MinPercent;