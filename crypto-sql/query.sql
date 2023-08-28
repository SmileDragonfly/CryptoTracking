-- name: InsertBUSDPrice :exec
INSERT INTO tblBUSDPrice ("price") VALUES ($1);
-- name: GetLastestBUSDPrice :one
SELECT "time", "price" FROM tblBUSDPrice ORDER BY "time" DESC LIMIT 1;
-- name: Get1MinAgoBUSDPrice :one
SELECT "time", "price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '61 seconds') ORDER BY "time" ASC LIMIT 1;
-- name: Get5MinAgoBUSDPrice :one
SELECT "time","price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '301 seconds') ORDER BY "time" ASC LIMIT 1;
-- name: Get10MinAgoBUSDPrice :one
SELECT "time","price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '601 seconds') ORDER BY "time" ASC LIMIT 1;
-- name: Get15MinAgoBUSDPrice :one
SELECT "time","price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '901 seconds') ORDER BY "time" ASC LIMIT 1;
-- name: Get30MinAgoBUSDPrice :one
SELECT "time","price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '1801 seconds') ORDER BY "time" ASC LIMIT 1;
-- name: Get60MinAgoBUSDPrice :one
SELECT "time","price" FROM tblBUSDPrice WHERE "time" > (NOW() - INTERVAL '3601 seconds') ORDER BY "time" ASC LIMIT 1;

-- name: Insert1MinBUSDPercent :exec
INSERT INTO tblBUSD1MinPercent ("prevtime", "symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4, $5);
-- name: Insert5MinBUSDPercent :exec
INSERT INTO tblBUSD5MinPercent ("prevtime", "symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4, $5);
-- name: Insert10MinBUSDPercent :exec
INSERT INTO tblBUSD10MinPercent ("prevtime", "symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4, $5);
-- name: Insert15MinBUSDPercent :exec
INSERT INTO tblBUSD15MinPercent ("prevtime", "symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4, $5);
-- name: Insert30MinBUSDPercent :exec
INSERT INTO tblBUSD30MinPercent ("prevtime", "symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4, $5);
-- name: Insert60MinBUSDPercent :exec
INSERT INTO tblBUSD60MinPercent ("prevtime", "symbol", "price", "prevprice", "percent") VALUES ($1, $2, $3, $4, $5);

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

-- name: GetAll1MinPercentDesc :many
SELECT * FROM tblBUSD1MinPercent ORDER BY "percent" DESC LIMIT $1;
-- name: GetAll5MinPercentDesc :many
SELECT * FROM tblBUSD5MinPercent ORDER BY "percent" DESC LIMIT $1;
-- name: GetAll10MinPercentDesc :many
SELECT * FROM tblBUSD10MinPercent ORDER BY "percent" DESC LIMIT $1;
-- name: GetAll15MinPercentDesc :many
SELECT * FROM tblBUSD15MinPercent ORDER BY "percent" DESC LIMIT $1;
-- name: GetAll30MinPercentDesc :many
SELECT * FROM tblBUSD30MinPercent ORDER BY "percent" DESC LIMIT $1;
-- name: GetAll60MinPercentDesc :many
SELECT * FROM tblBUSD60MinPercent ORDER BY "percent" DESC LIMIT $1;

-- name: GetAll1MinPercentAsc :many
SELECT * FROM tblBUSD1MinPercent ORDER BY "percent" ASC LIMIT $1;
-- name: GetAll5MinPercentAsc :many
SELECT * FROM tblBUSD5MinPercent ORDER BY "percent" ASC LIMIT $1;
-- name: GetAll10MinPercentAsc :many
SELECT * FROM tblBUSD10MinPercent ORDER BY "percent" ASC LIMIT $1;
-- name: GetAll15MinPercentAsc :many
SELECT * FROM tblBUSD15MinPercent ORDER BY "percent" ASC LIMIT $1;
-- name: GetAll30MinPercentAsc :many
SELECT * FROM tblBUSD30MinPercent ORDER BY "percent" ASC LIMIT $1;
-- name: GetAll60MinPercentAsc :many
SELECT * FROM tblBUSD60MinPercent ORDER BY "percent" ASC LIMIT $1;

-- name: DeleteWasteBUSDPrice :exec
DELETE FROM tblBUSDPrice WHERE "time" < (NOW() - INTERVAL '7200 seconds');

-- name: InsertTopCoinHistory :exec
INSERT INTO tblBUSDTopCoinHistory (topCoin) VALUES ($1);
-- name: GetTopCoinHistory :many
SELECT * FROM tblBUSDTopCoinHistory ORDER BY "time" DESC LIMIT $1;
-- name: DeleteTopCoinHistory :exec
DELETE FROM tblBUSDTopCoinHistory WHERE "time" < (NOW() - INTERVAL '1 days');

