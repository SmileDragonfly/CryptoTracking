package main

import (
	"CryptoTrackingSql/sqlc"
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
)

func TestCalculatePercent(t *testing.T) {
	err := loadConfig(".")
	if err != nil {
		log.Fatalf(err.Error())
	}
	GStrConn = fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	defer conn.Close()
	query := sqlc.New(conn)
	strPrice, err := query.GetLastestBUSDPrice(context.Background())
	strPrice1MinAgo, err := query.Get60MinAgoBUSDPrice(context.Background())
	percent, err := CalculatePercent(strPrice.String, strPrice1MinAgo.String)
	if err != nil {
		t.Error(err.Error())
	}
	if len(percent) == 0 {
		t.Error(getCurrentFuncname(), "Expected > 0 but got 0")
	}
}
