package main

import (
	"context"
	"cryptoagent/logger"
	"cryptosql/sqlc"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"sort"
	"time"
)

// run will be called by Start() so business logic goes here
func main() {
	err := logger.NewLogger("./config/logcfg.json")
	if err != nil {
		panic(err)
	}
	logger.Info("Start logger succesfully")
	// Begin main program
	err = loadConfig("./config/")
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	GStrConn = fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	defer conn.Close()
	queries := sqlc.New(conn)
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
					logger.Error(err.Error())
				} else {
					// Calculate percent
					//if err != nil {
					//	logger.Error(err.Error())
					//	conn.Close()
					//	continue
					//}
					//tx, err := conn.Begin()
					//if err != nil {
					//	logger.Error(err.Error())
					//	conn.Close()
					//	continue
					//}
					//query := queries.WithTx(tx)
					// 2.Get price 1,5,10,15,30,60 min ago
					strPriceLastest, err := queries.GetLastestBUSDPrice(context.Background())
					strPrice1MinAgo, err := queries.Get1MinAgoBUSDPrice(context.Background())
					strPrice5MinAgo, err := queries.Get5MinAgoBUSDPrice(context.Background())
					strPrice10MinAgo, err := queries.Get10MinAgoBUSDPrice(context.Background())
					strPrice15MinAgo, err := queries.Get15MinAgoBUSDPrice(context.Background())
					strPrice30MinAgo, err := queries.Get30MinAgoBUSDPrice(context.Background())
					strPrice60MinAgo, err := queries.Get60MinAgoBUSDPrice(context.Background())
					chanPercentLastest := make(chan []TPricePercent)
					chanPercent1Min := make(chan []TPricePercent)
					chanPercent5Min := make(chan []TPricePercent)
					chanPercent10Min := make(chan []TPricePercent)
					chanPercent15Min := make(chan []TPricePercent)
					chanPercent30Min := make(chan []TPricePercent)
					chanPercent60Min := make(chan []TPricePercent)
					logger.Info("Begin calculate percent")
					go CalculatePercentRountine(strPrice, strPriceLastest.Price.String, chanPercentLastest)
					go CalculatePercentRountine(strPrice, strPrice1MinAgo.Price.String, chanPercent1Min)
					go CalculatePercentRountine(strPrice, strPrice5MinAgo.Price.String, chanPercent5Min)
					go CalculatePercentRountine(strPrice, strPrice10MinAgo.Price.String, chanPercent10Min)
					go CalculatePercentRountine(strPrice, strPrice15MinAgo.Price.String, chanPercent15Min)
					go CalculatePercentRountine(strPrice, strPrice30MinAgo.Price.String, chanPercent30Min)
					go CalculatePercentRountine(strPrice, strPrice60MinAgo.Price.String, chanPercent60Min)
					arrPercentLastest := <-chanPercentLastest
					arrPercent1Min := <-chanPercent1Min
					arrPercent5Min := <-chanPercent5Min
					arrPercent10Min := <-chanPercent10Min
					arrPercent15Min := <-chanPercent15Min
					arrPercent30Min := <-chanPercent30Min
					arrPercent60Min := <-chanPercent60Min
					logger.Info("End calculate percent")
					if arrPercent1Min == nil {
						logger.Warning("CalculatePercent 1 Min NUll Data")
					} else {
						queries.Delete1MinBUSDPercent(context.Background())
						//logger.Debugf("Insert1MinBUSDPercent begin")
						//for _, v := range arrPercent1Min {
						//	queries.Insert1MinBUSDPercent(context.Background(),
						//		sqlc.Insert1MinBUSDPercentParams{strPrice1MinAgo.Time,
						//			sql.NullString{v.Symbol, true},
						//			sql.NullFloat64{v.Price, true},
						//			sql.NullFloat64{v.PrevPrice, true},
						//			sql.NullFloat64{v.Percent, true}})
						//}
						//logger.Debugf("Insert1MinBUSDPercent end")
						logger.Debugf("Insert1MinBUSDPercent begin insert all")
						rows := []sqlc.Insert1MinBUSDPercentParams{}
						for _, it := range arrPercent1Min {
							row := sqlc.Insert1MinBUSDPercentParams{
								Prevtime:  strPrice1MinAgo.Time,
								Symbol:    sql.NullString{it.Symbol, true},
								Price:     sql.NullFloat64{it.Price, true},
								Prevprice: sql.NullFloat64{it.PrevPrice, true},
								Percent:   sql.NullFloat64{it.Percent, true},
							}
							rows = append(rows, row)
						}
						queries.Insert1MinBUSDPercentRows(context.Background(), rows)
						logger.Debugf("Insert1MinBUSDPercent end insert all")
					}
					if arrPercent5Min == nil {
						logger.Warning("CalculatePercent 5 Min NUll Data")
					} else {
						logger.Debugf("Insert5MinBUSDPercent begin")
						queries.Delete5MinBUSDPercent(context.Background())
						for _, v := range arrPercent5Min {
							queries.Insert5MinBUSDPercent(context.Background(),
								sqlc.Insert5MinBUSDPercentParams{strPrice5MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
						logger.Debugf("Insert30MinBUSDPercent end")
					}
					if arrPercent10Min == nil {
						logger.Warning("CalculatePercent 10 Min NUll Data")
					} else {
						logger.Debugf("Insert10MinBUSDPercent begin")
						queries.Delete10MinBUSDPercent(context.Background())
						for _, v := range arrPercent10Min {
							queries.Insert10MinBUSDPercent(context.Background(),
								sqlc.Insert10MinBUSDPercentParams{strPrice10MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
						logger.Debugf("Insert30MinBUSDPercent end")
					}
					if arrPercent15Min == nil {
						logger.Warning("CalculatePercent 15 Min NUll Data")
					} else {
						logger.Debugf("Insert15MinBUSDPercent begin")
						queries.Delete15MinBUSDPercent(context.Background())
						for _, v := range arrPercent15Min {
							queries.Insert15MinBUSDPercent(context.Background(),
								sqlc.Insert15MinBUSDPercentParams{strPrice15MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
						logger.Debugf("Insert30MinBUSDPercent end")
					}
					if arrPercent30Min == nil {
						logger.Warning("CalculatePercent 30 Min NUll Data")
					} else {
						logger.Debugf("Insert30MinBUSDPercent begin")
						queries.Delete30MinBUSDPercent(context.Background())
						for _, v := range arrPercent30Min {
							queries.Insert30MinBUSDPercent(context.Background(),
								sqlc.Insert30MinBUSDPercentParams{strPrice30MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
						logger.Debugf("Insert30MinBUSDPercent end")
					}
					if arrPercent60Min == nil {
						logger.Warning("CalculatePercent 60 Min NUll Data")
					} else {
						logger.Debugf("Insert60MinBUSDPercent begin")
						queries.Delete60MinBUSDPercent(context.Background())
						for _, v := range arrPercent60Min {
							queries.Insert60MinBUSDPercent(context.Background(),
								sqlc.Insert60MinBUSDPercentParams{strPrice60MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
						logger.Debugf("Insert60MinBUSDPercent end")
					}
					if arrPercentLastest == nil {
						logger.Warning("CalculatePercent Lastest NUll Data")
					} else {
						logger.Debugf("InsertTopCoinHistory begin")
						sort.Slice(arrPercentLastest, func(i, j int) bool {
							return arrPercentLastest[i].Percent > arrPercentLastest[j].Percent
						})
						arrPercentLastest = arrPercentLastest[:GConfig.NumberOfTopcoint]
						arrTopCoin := []TTopCoin{}
						for _, v := range arrPercentLastest {
							it := TTopCoin{}
							it.Symbol = v.Symbol
							it.Percent = v.Percent
							arrTopCoin = append(arrTopCoin, it)
						}
						// Convert to json
						strTopCoin, err := json.Marshal(&arrTopCoin)
						if err != nil {
							logger.Error("Convert arrTopCoin to string failed:", err)
						}
						err = queries.InsertTopCoinHistory(context.Background(), sql.NullString{
							String: string(strTopCoin),
							Valid:  true,
						})
						if err != nil {
							logger.Error("Insert to TopCoinHistory failed:", err)
						}
						logger.Debugf("InsertTopCoinHistory end")
					}
					// Insert to DB
					logger.Debugf("InsertBUSDPrice")
					err = queries.InsertBUSDPrice(context.Background(), sql.NullString{strPrice, true})
					if err != nil {
						logger.Error(err.Error())
					}
					// Delete waste data
					err = queries.DeleteWasteBUSDPrice(context.Background())
					if err != nil {
						logger.Error(err.Error())
					}
					// Delete waste data
					err = queries.DeleteTopCoinHistory(context.Background())
					if err != nil {
						logger.Error(err.Error())
					}
					//tx.Commit()
				}
			}
		}
	}()
	<-tickerDone
}
