package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var GStrConn string

func main() {
	// Init log file
	f, err := os.OpenFile("CryptoTrackingAgent.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Start CryptoTrackingAgent")
	// Begin main program
	var api BinanceAPI
	loadConfig(".")
	strPrice := api.getAllPrice("BUSD$")
	// Insert to DB
	GStrConn = fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)
	err = insertTblBUSDPrice(strPrice, GStrConn, GConfig.DBDriver)
	if err != nil {
		log.Println(err.Error())
	}
}
