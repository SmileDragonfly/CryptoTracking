package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type BinanceAPI struct{}
type TTickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (receiver BinanceAPI) getAllPrice(symReg string) string {
	link := "https://api.binance.com/api/v3/ticker/price"
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(getCurrentFuncname(), err.Error())
	}
	respBody, err := io.ReadAll(resp.Body)
	//respStr := string(respBody)
	//fmt.Println(getCurrentFuncname(), respStr)
	var arrJson []TTickerPrice
	err = json.Unmarshal(respBody, &arrJson)
	if err != nil {
		fmt.Println(getCurrentFuncname(), err.Error())
	}
	// Get all symbol match with symReg
	var arrRet []TTickerPrice
	for _, v := range arrJson {
		bMatch, _ := regexp.MatchString(symReg, v.Symbol)
		if bMatch {
			arrRet = append(arrRet, v)
		}
	}
	// Marshal arrRet to strRet
	var byteRet []byte
	byteRet, err = json.Marshal(arrRet)
	if err != nil {
		fmt.Println(getCurrentFuncname(), err.Error())
	}
	return string(byteRet)
}
