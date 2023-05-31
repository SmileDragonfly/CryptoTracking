package main

import (
	"context"
	"cryptoagent/logger"
	"cryptosql/sqlc"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"sort"
	"time"
)

// run will be called by Start() so business logic goes here
func main() {
	// Parse log config
	byteCfg, err := os.ReadFile("./config/logcfg.json")
	if err != nil {
		panic(err)
	}
	var logCfg logger.LoggerConfig
	err = json.Unmarshal(byteCfg, &logCfg)
	if err != nil {
		panic(err)
	}
	err = logger.NewLogger(logCfg)
	if err != nil {
		panic(err)
	}
	logger.Logger.Info("Start logger succesfully")
	// Begin main program
	err = loadConfig("./config/")
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
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
					logger.Logger.Error(err.Error())
				} else {
					// Calculate percent
					// 1.Open DB
					conn, err := sql.Open(GConfig.DBDriver, GStrConn)
					if err != nil {
						logger.Logger.Error(err.Error())
						conn.Close()
						continue
					}
					tx, err := conn.Begin()
					if err != nil {
						logger.Logger.Error(err.Error())
						conn.Close()
						continue
					}
					query := sqlc.New(conn)
					query = query.WithTx(tx)
					// 2.Get price 1,5,10,15,30,60 min ago
					strPriceLastest, err := query.GetLastestBUSDPrice(context.Background())
					strPrice1MinAgo, err := query.Get1MinAgoBUSDPrice(context.Background())
					strPrice5MinAgo, err := query.Get5MinAgoBUSDPrice(context.Background())
					strPrice10MinAgo, err := query.Get10MinAgoBUSDPrice(context.Background())
					strPrice15MinAgo, err := query.Get15MinAgoBUSDPrice(context.Background())
					strPrice30MinAgo, err := query.Get30MinAgoBUSDPrice(context.Background())
					strPrice60MinAgo, err := query.Get60MinAgoBUSDPrice(context.Background())
					chanPercentLastest := make(chan []TPricePercent)
					chanPercent1Min := make(chan []TPricePercent)
					chanPercent5Min := make(chan []TPricePercent)
					chanPercent10Min := make(chan []TPricePercent)
					chanPercent15Min := make(chan []TPricePercent)
					chanPercent30Min := make(chan []TPricePercent)
					chanPercent60Min := make(chan []TPricePercent)
					logger.Logger.Info("Begin calculate percent")
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
					logger.Logger.Info("End calculate percent")
					if arrPercent1Min == nil {
						logger.Logger.Warning("CalculatePercent 1 Min NUll Data")
					} else {
						query.Delete1MinBUSDPercent(context.Background())
						for _, v := range arrPercent1Min {
							query.Insert1MinBUSDPercent(context.Background(),
								sqlc.Insert1MinBUSDPercentParams{strPrice1MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercent5Min == nil {
						logger.Logger.Warning("CalculatePercent 5 Min NUll Data")
					} else {
						query.Delete5MinBUSDPercent(context.Background())
						for _, v := range arrPercent5Min {
							query.Insert5MinBUSDPercent(context.Background(),
								sqlc.Insert5MinBUSDPercentParams{strPrice5MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercent10Min == nil {
						logger.Logger.Warning("CalculatePercent 10 Min NUll Data")
					} else {
						query.Delete10MinBUSDPercent(context.Background())
						for _, v := range arrPercent10Min {
							query.Insert10MinBUSDPercent(context.Background(),
								sqlc.Insert10MinBUSDPercentParams{strPrice10MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercent15Min == nil {
						logger.Logger.Warning("CalculatePercent 15 Min NUll Data")
					} else {
						query.Delete15MinBUSDPercent(context.Background())
						for _, v := range arrPercent15Min {
							query.Insert15MinBUSDPercent(context.Background(),
								sqlc.Insert15MinBUSDPercentParams{strPrice15MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercent30Min == nil {
						logger.Logger.Warning("CalculatePercent 30 Min NUll Data")
					} else {
						query.Delete30MinBUSDPercent(context.Background())
						for _, v := range arrPercent30Min {
							query.Insert30MinBUSDPercent(context.Background(),
								sqlc.Insert30MinBUSDPercentParams{strPrice30MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercent60Min == nil {
						logger.Logger.Warning("CalculatePercent 60 Min NUll Data")
					} else {
						query.Delete60MinBUSDPercent(context.Background())
						for _, v := range arrPercent60Min {
							query.Insert60MinBUSDPercent(context.Background(),
								sqlc.Insert60MinBUSDPercentParams{strPrice60MinAgo.Time,
									sql.NullString{v.Symbol, true},
									sql.NullFloat64{v.Price, true},
									sql.NullFloat64{v.PrevPrice, true},
									sql.NullFloat64{v.Percent, true}})
						}
					}
					if arrPercentLastest == nil {
						logger.Logger.Warning("CalculatePercent Lastest NUll Data")
					} else {
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
							logger.Logger.Error("Convert arrTopCoin to string failed:", err)
						}
						err = query.InsertTopCoinHistory(context.Background(), sql.NullString{
							String: string(strTopCoin),
							Valid:  true,
						})
						if err != nil {
							logger.Logger.Error("Insert to TopCoinHistory failed:", err)
						}
					}
					// Insert to DB
					err = query.InsertBUSDPrice(context.Background(), sql.NullString{strPrice, true})
					if err != nil {
						logger.Logger.Error(err.Error())
					}
					// Delete waste data
					err = query.DeleteWasteBUSDPrice(context.Background())
					if err != nil {
						logger.Logger.Error(err.Error())
					}
					// Delete waste data
					err = query.DeleteTopCoinHistory(context.Background())
					if err != nil {
						logger.Logger.Error(err.Error())
					}
					tx.Commit()
					conn.Close()
				}
			}
		}
	}()
	<-tickerDone
}
