package main

import (
	"CryptoTrackingSql/sqlc"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
)

type TPricePercent struct {
	Symbol    string  `json:"symbol"`
	Price     string  `json:"price"`
	PrevPrice string  `json:"prevPrice"`
	Percent   float64 `json:"percent"`
}

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

func CalculatePercent(strPrice string, stringPrePrice string) ([]TPricePercent, error) {
	var arrPrice []TTickerPrice
	var arrPrevPrice []TTickerPrice
	var arrPricePercent []TPricePercent
	json.Unmarshal([]byte(strPrice), &arrPrice)
	json.Unmarshal([]byte(stringPrePrice), &arrPrevPrice)
	for _, v := range arrPrice {
		for _, vPre := range arrPrevPrice {
			if v.Symbol == vPre.Symbol {
				var pricePercent TPricePercent
				pricePercent.Symbol = v.Symbol
				pricePercent.Price = v.Price
				pricePercent.PrevPrice = vPre.Price
				fPrice, err := strconv.ParseFloat(v.Price, 32)
				if err != nil {
					return nil, err
				}
				fPrePrice, err := strconv.ParseFloat(vPre.Price, 32)
				if err != nil {
					return nil, err
				}
				pricePercent.Percent = (fPrice - fPrePrice) / fPrePrice
				arrPricePercent = append(arrPricePercent, pricePercent)
				continue
			}
		}
	}
	return arrPricePercent, nil
}
