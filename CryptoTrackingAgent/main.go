package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
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
	err = loadConfig(".")
	if err != nil {
		log.Fatalf(err.Error())
	}
	GStrConn = fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)

	// Setup a ticker
	ticker := time.NewTicker(time.Duration(GConfig.TickerGetPrice) * time.Second)
	tickerDone := make(chan bool)
	go func() {
		for {
			select {
			case <-tickerDone:
				return
			case <-ticker.C:
				strPrice, err := api.getAllPrice("BUSD$")
				if err != nil {
					log.Println(err.Error())
				} else {
					// Insert to DB
					err = insertTblBUSDPrice(strPrice, GStrConn, GConfig.DBDriver)
					if err != nil {
						log.Println(err.Error())
					}
					log.Print(getCurrentFuncname())
				}
			}
		}
	}()
	<-tickerDone
}
