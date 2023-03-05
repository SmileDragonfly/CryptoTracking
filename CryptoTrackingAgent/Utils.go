package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
)

type TPricePercent struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	PrevPrice float64 `json:"prevPrice"`
	Percent   float64 `json:"percent"`
}

func getCurrentFuncname() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}

func CalculatePercent(strPrice string, stringPrePrice string) ([]TPricePercent, error) {
	var arrPrice []TTickerPrice
	var arrPrevPrice []TTickerPrice
	var arrPricePercent []TPricePercent
	err := json.Unmarshal([]byte(strPrice), &arrPrice)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(stringPrePrice), &arrPrevPrice)
	if err != nil {
		return nil, err
	}
	for _, v := range arrPrice {
		for _, vPre := range arrPrevPrice {
			if v.Symbol == vPre.Symbol {
				var pricePercent TPricePercent
				pricePercent.Symbol = v.Symbol

				fPrice, err := strconv.ParseFloat(v.Price, 32)
				if err != nil {
					return nil, err
				}
				fPrePrice, err := strconv.ParseFloat(vPre.Price, 32)
				if err != nil {
					return nil, err
				}
				pricePercent.Price = fPrice
				pricePercent.PrevPrice = fPrePrice
				pricePercent.Percent = (fPrice - fPrePrice) / fPrePrice
				arrPricePercent = append(arrPricePercent, pricePercent)
				continue
			}
		}
	}
	return arrPricePercent, nil
}
