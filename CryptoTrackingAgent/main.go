package main

import (
	"CryptoTrackingSql/sqlc"
	"context"
	"database/sql"
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
	log.Println("----------Start CryptoTrackingAgent----------")

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
					// Calculate percent
					// 1.Open DB
					conn, err := sql.Open(GConfig.DBDriver, GStrConn)
					if err != nil {
						log.Println(err.Error())
					}
					query := sqlc.New(conn)
					// 2.Get price 1,5,10,15,30,60 min ago
					strPrice1MinAgo, err := query.Get1MinAgoBUSDPrice(context.Background())
					//strPrice5MinAgo, err := query.Get1MinAgoBUSDPrice(context.Background())
					//strPrice10MinAgo, err := query.Get1MinAgoBUSDPrice(context.Background())
					//strPrice15MinAgo, err := query.Get1MinAgoBUSDPrice(context.Background())
					//strPrice30MinAgo, err := query.Get1MinAgoBUSDPrice(context.Background())
					//strPrice60MinAgo, err := query.Get1MinAgoBUSDPrice(context.Background())
					arrPercent1Min, err := CalculatePercent(strPrice, strPrice1MinAgo.String)
					if err != nil {
						log.Println("CalculatePercent 1 Min", err.Error())
					} else {
						for _, v := range arrPercent1Min {
							query.Insert1MinBUSDPercent(context.Background(),
								sqlc.Insert1MinBUSDPercentParams{sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}

					err = query.InsertBUSDPrice(context.Background(), sql.NullString{strPrice, true})
					// Insert to DB
					if err != nil {
						log.Println(err.Error())
					}
					conn.Close()
				}
			}
		}
	}()
	<-tickerDone
}
