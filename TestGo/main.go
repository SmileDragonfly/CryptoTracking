package main

import (
	"encoding/json"
	"fmt"
)

type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func main() {
	jsonString := `[{"symbol":"ETHBTC","price":"0.07023000"},{"symbol":"LTCBTC","price":"0.00412800"}]`

	var arr []SymbolPrice
	err := json.Unmarshal([]byte(jsonString), &arr)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Here is a trick. Replace nil with an empty slice.
	fmt.Println(arr[0])
}
