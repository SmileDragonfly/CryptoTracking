package main

import (
	"CryptoTrackingSql/sqlc"
	"context"
	"database/sql"
	"fmt"
	log "github.com/jeanphorn/log4go"
	_ "github.com/lib/pq"
	"time"
)

// run will be called by Start() so business logic goes here
func main() {
	// Init log file
	// load config file, it's optional
	// or log.LoadConfiguration("./example.json", "json")
	// config file could be json or xml
	log.LoadConfiguration("./log4go.json")
	log.Info("----------Start CryptoTrackingAgent----------")
	defer log.Close()
	// Begin main program
	err := loadConfig(".")
	if err != nil {
		log.Error(err.Error())
	}
	GStrConn = fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)
	// Setup a ticker
	var api BinanceAPI
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
					log.Error(err.Error())
				} else {
					// Calculate percent
					// 1.Open DB
					conn, err := sql.Open(GConfig.DBDriver, GStrConn)
					if err != nil {
						log.Error(err.Error())
						conn.Close()
						continue
					}
					tx, err := conn.Begin()
					if err != nil {
						log.Error(err.Error())
						conn.Close()
						continue
					}
					query := sqlc.New(conn)
					query = query.WithTx(tx)
					// 2.Get price 1,5,10,15,30,60 min ago
					strPrice1MinAgo, err := query.Get1MinAgoBUSDPrice(context.Background())
					strPrice5MinAgo, err := query.Get5MinAgoBUSDPrice(context.Background())
					strPrice10MinAgo, err := query.Get10MinAgoBUSDPrice(context.Background())
					strPrice15MinAgo, err := query.Get15MinAgoBUSDPrice(context.Background())
					strPrice30MinAgo, err := query.Get30MinAgoBUSDPrice(context.Background())
					strPrice60MinAgo, err := query.Get60MinAgoBUSDPrice(context.Background())
					chanPercent1Min := make(chan []TPricePercent)
					chanPercent5Min := make(chan []TPricePercent)
					chanPercent10Min := make(chan []TPricePercent)
					chanPercent15Min := make(chan []TPricePercent)
					chanPercent30Min := make(chan []TPricePercent)
					chanPercent60Min := make(chan []TPricePercent)
					log.Info("Begin calculate percent")
					go CalculatePercentRountine(strPrice, strPrice1MinAgo.String, chanPercent1Min)
					go CalculatePercentRountine(strPrice, strPrice5MinAgo.String, chanPercent5Min)
					go CalculatePercentRountine(strPrice, strPrice10MinAgo.String, chanPercent10Min)
					go CalculatePercentRountine(strPrice, strPrice15MinAgo.String, chanPercent15Min)
					go CalculatePercentRountine(strPrice, strPrice30MinAgo.String, chanPercent30Min)
					go CalculatePercentRountine(strPrice, strPrice60MinAgo.String, chanPercent60Min)
					arrPercent1Min := <-chanPercent1Min
					arrPercent5Min := <-chanPercent5Min
					arrPercent10Min := <-chanPercent10Min
					arrPercent15Min := <-chanPercent15Min
					arrPercent30Min := <-chanPercent30Min
					arrPercent60Min := <-chanPercent60Min
					log.Info("End calculate percent")
					if arrPercent1Min == nil {
						log.Warn("CalculatePercent 1 Min NUll Data")
					} else {
						query.Delete1MinBUSDPercent(context.Background())
						for _, v := range arrPercent1Min {
							query.Insert1MinBUSDPercent(context.Background(),
								sqlc.Insert1MinBUSDPercentParams{sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercent5Min == nil {
						log.Warn("CalculatePercent 5 Min NUll Data")
					} else {
						query.Delete5MinBUSDPercent(context.Background())
						for _, v := range arrPercent5Min {
							query.Insert5MinBUSDPercent(context.Background(),
								sqlc.Insert5MinBUSDPercentParams{sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercent10Min == nil {
						log.Warn("CalculatePercent 10 Min NUll Data")
					} else {
						query.Delete10MinBUSDPercent(context.Background())
						for _, v := range arrPercent10Min {
							query.Insert10MinBUSDPercent(context.Background(),
								sqlc.Insert10MinBUSDPercentParams{sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercent15Min == nil {
						log.Warn("CalculatePercent 15 Min NUll Data")
					} else {
						query.Delete15MinBUSDPercent(context.Background())
						for _, v := range arrPercent15Min {
							query.Insert15MinBUSDPercent(context.Background(),
								sqlc.Insert15MinBUSDPercentParams{sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercent30Min == nil {
						log.Warn("CalculatePercent 30 Min NUll Data")
					} else {
						query.Delete30MinBUSDPercent(context.Background())
						for _, v := range arrPercent30Min {
							query.Insert30MinBUSDPercent(context.Background(),
								sqlc.Insert30MinBUSDPercentParams{sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercent60Min == nil {
						log.Warn("CalculatePercent 60 Min NUll Data")
					} else {
						query.Delete60MinBUSDPercent(context.Background())
						for _, v := range arrPercent60Min {
							query.Insert60MinBUSDPercent(context.Background(),
								sqlc.Insert60MinBUSDPercentParams{sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}

					err = query.InsertBUSDPrice(context.Background(), sql.NullString{strPrice, true})
					// Insert to DB
					if err != nil {
						log.Error(err.Error())
					}
					tx.Commit()
					conn.Close()
				}
			}
		}
	}()
	<-tickerDone
}
