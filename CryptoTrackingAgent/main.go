package main

import (
	"CryptoTrackingSql/sqlc"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	var api BinanceAPI
	loadConfig(".")
	strPrice := api.getAllPrice("BUSD$")
	// Insert to DB
	strConn := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)
	conn, err := sql.Open(GConfig.DBDriver, strConn)
	if err != nil {
		fmt.Print(getCurrentFuncname(), err.Error())
	}
	query := sqlc.New(conn)
	err = query.InsertBUSDPrice(context.Background(), sql.NullString{strPrice, true})
	if err != nil {
		fmt.Print(getCurrentFuncname(), err.Error())
	}
}
