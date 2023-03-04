package main

import (
	"CryptoTrackingSql/sqlc"
	"context"
	"database/sql"
	"fmt"
	"runtime"
)

func getCurrentFuncname() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}

func insertTblBUSDPrice(strPrice string, strDBConn string, strDBDriver string) error {
	conn, err := sql.Open(strDBDriver, strDBConn)
	defer conn.Close()
	if err != nil {
		return err
	}
	query := sqlc.New(conn)
	err = query.InsertBUSDPrice(context.Background(), sql.NullString{strPrice, true})
	return err
}
