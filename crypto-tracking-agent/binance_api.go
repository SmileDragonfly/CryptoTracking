package main

import (
	"cryptoagent/logger"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

type BinanceAPI struct{}
type TTickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (b BinanceAPI) getAllPrice(symReg string) (string, error) {
	link := "https://api.binance.com/api/v3/ticker/price"
	resp, err := http.Get(link)
	if err != nil {
		logger.Logger.Error(err.Error())
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.Error(err.Error())
		return "", err
	}
	//respStr := string(respBody)
	//fmt.Println(getCurrentFuncname(), respStr)
	var arrJson []TTickerPrice
	err = json.Unmarshal(respBody, &arrJson)
	if err != nil {
		logger.Logger.Error(err.Error())
		return "", err
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
		logger.Logger.Error(err.Error())
		return "", err
	}
	return string(byteRet), err
}
