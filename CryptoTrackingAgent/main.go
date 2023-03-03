package main

import "fmt"

func main() {
	var api BinanceAPI
	strPrice := api.getAllPrice("BUSD$")
	fmt.Println(getCurrentFuncname(), strPrice)
}
