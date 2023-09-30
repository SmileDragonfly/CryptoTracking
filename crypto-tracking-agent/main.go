package main

import (
	"context"
	"cryptoagent/logger"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/SmileDragonfly/go-lib/crypto-sql/sqlc"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	GStrConn = fmt.Sprintf("port=%d host=%s user=%s password=%s database=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	err = conn.Ping()
	if err != nil {
		strConn := fmt.Sprintf("port=%d host=%s user=%s password=%s sslmode=disable",
			GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password)
		defaultConn, err := sql.Open(GConfig.DBDriver, strConn)
		if err != nil {
			panic(err)
		}
		// Check db is exist
		query := "SELECT datname FROM pg_database WHERE datname='%s';"
		query = fmt.Sprintf(query, GConfig.DBName)
		var dbname string
		if err := defaultConn.QueryRow(query).Scan(&dbname); err != nil {
			if err != sql.ErrNoRows {
				panic(err)
			}
		}
		if dbname == "" {
			// Create database
			query := "CREATE DATABASE %s;"
			query = fmt.Sprintf(query, GConfig.DBName)
			_, err := defaultConn.Exec(query)
			if err != nil {
				panic(err)
			}
		}
		defaultConn.Close()
		if err := conn.Ping(); err != nil {
			panic(err)
		}
	}
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://./db", "crypto", driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		panic(err)
	}
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
						logger.Debugf("Begin Insert1MinBUSDPercentRows")
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
						err := queries.Insert1MinBUSDPercentRows(context.Background(), rows)
						if err != nil {
							logger.Errorf("Insert1MinBUSDPercentRows: %s", err)
						}
						logger.Debugf("End Insert1MinBUSDPercentRows")
					}

					if arrPercent5Min == nil {
						logger.Warning("CalculatePercent 5 Min NUll Data")
					} else {
						queries.Delete5MinBUSDPercent(context.Background())
						logger.Debugf("Begin Insert5MinBUSDPercentRows")
						rows := []sqlc.Insert5MinBUSDPercentParams{}
						for _, it := range arrPercent5Min {
							row := sqlc.Insert5MinBUSDPercentParams{
								Prevtime:  strPrice5MinAgo.Time,
								Symbol:    sql.NullString{it.Symbol, true},
								Price:     sql.NullFloat64{it.Price, true},
								Prevprice: sql.NullFloat64{it.PrevPrice, true},
								Percent:   sql.NullFloat64{it.Percent, true},
							}
							rows = append(rows, row)
						}
						err := queries.Insert5MinBUSDPercentRows(context.Background(), rows)
						if err != nil {
							logger.Errorf("Insert5MinBUSDPercentRows: %s", err)
						}
						logger.Debugf("End Insert5MinBUSDPercentRows")
					}

					if arrPercent10Min == nil {
						logger.Warning("CalculatePercent 10 Min NUll Data")
					} else {
						queries.Delete10MinBUSDPercent(context.Background())
						logger.Debugf("Begin Insert10MinBUSDPercentRows")
						rows := []sqlc.Insert10MinBUSDPercentParams{}
						for _, it := range arrPercent10Min {
							row := sqlc.Insert10MinBUSDPercentParams{
								Prevtime:  strPrice10MinAgo.Time,
								Symbol:    sql.NullString{it.Symbol, true},
								Price:     sql.NullFloat64{it.Price, true},
								Prevprice: sql.NullFloat64{it.PrevPrice, true},
								Percent:   sql.NullFloat64{it.Percent, true},
							}
							rows = append(rows, row)
						}
						err := queries.Insert10MinBUSDPercentRows(context.Background(), rows)
						if err != nil {
							logger.Errorf("Insert10MinBUSDPercentRows: %s", err)
						}
						logger.Debugf("End Insert10MinBUSDPercentRows")
					}

					if arrPercent15Min == nil {
						logger.Warning("CalculatePercent 15 Min NUll Data")
					} else {
						queries.Delete15MinBUSDPercent(context.Background())
						logger.Debugf("Begin Insert15MinBUSDPercentRows")
						rows := []sqlc.Insert15MinBUSDPercentParams{}
						for _, it := range arrPercent10Min {
							row := sqlc.Insert15MinBUSDPercentParams{
								Prevtime:  strPrice15MinAgo.Time,
								Symbol:    sql.NullString{it.Symbol, true},
								Price:     sql.NullFloat64{it.Price, true},
								Prevprice: sql.NullFloat64{it.PrevPrice, true},
								Percent:   sql.NullFloat64{it.Percent, true},
							}
							rows = append(rows, row)
						}
						err := queries.Insert15MinBUSDPercentRows(context.Background(), rows)
						if err != nil {
							logger.Errorf("Insert15MinBUSDPercentRows: %s", err)
						}
						logger.Debugf("End Insert15MinBUSDPercentRows")
					}

					if arrPercent30Min == nil {
						logger.Warning("CalculatePercent 30 Min NUll Data")
					} else {
						queries.Delete30MinBUSDPercent(context.Background())
						logger.Debugf("Begin Insert30MinBUSDPercentRows")
						rows := []sqlc.Insert30MinBUSDPercentParams{}
						for _, it := range arrPercent30Min {
							row := sqlc.Insert30MinBUSDPercentParams{
								Prevtime:  strPrice30MinAgo.Time,
								Symbol:    sql.NullString{it.Symbol, true},
								Price:     sql.NullFloat64{it.Price, true},
								Prevprice: sql.NullFloat64{it.PrevPrice, true},
								Percent:   sql.NullFloat64{it.Percent, true},
							}
							rows = append(rows, row)
						}
						err := queries.Insert30MinBUSDPercentRows(context.Background(), rows)
						if err != nil {
							logger.Errorf("Insert30MinBUSDPercentRows: %s", err)
						}
						logger.Debugf("End Insert30MinBUSDPercentRows")
					}

					if arrPercent60Min == nil {
						logger.Warning("CalculatePercent 60 Min NUll Data")
					} else {
						queries.Delete60MinBUSDPercent(context.Background())
						logger.Debugf("Begin Insert60MinBUSDPercentRows")
						rows := []sqlc.Insert60MinBUSDPercentParams{}
						for _, it := range arrPercent60Min {
							row := sqlc.Insert60MinBUSDPercentParams{
								Prevtime:  strPrice60MinAgo.Time,
								Symbol:    sql.NullString{it.Symbol, true},
								Price:     sql.NullFloat64{it.Price, true},
								Prevprice: sql.NullFloat64{it.PrevPrice, true},
								Percent:   sql.NullFloat64{it.Percent, true},
							}
							rows = append(rows, row)
						}
						err := queries.Insert60MinBUSDPercentRows(context.Background(), rows)
						if err != nil {
							logger.Errorf("Insert60MinBUSDPercentRows: %s", err)
						}
						logger.Debugf("End Insert60MinBUSDPercentRows")
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
							logger.Errorf("InsertTopCoinHistory:%s", err)
						}
						logger.Debugf("InsertTopCoinHistory end")
					}
					// Insert to DB
					logger.Debugf("InsertBUSDPrice")
					err = queries.InsertBUSDPrice(context.Background(), sql.NullString{strPrice, true})
					if err != nil {
						logger.Error("InsertBUSDPrice: %s", err)
					}
					// Delete waste data
					err = queries.DeleteWasteBUSDPrice(context.Background())
					if err != nil {
						logger.Error("DeleteWasteBUSDPrice: %s", err)
					}
					// Delete waste data
					err = queries.DeleteTopCoinHistory(context.Background())
					if err != nil {
						logger.Error("DeleteTopCoinHistory: %s", err)
					}
					//tx.Commit()
				}
			}
		}
	}()
	<-tickerDone
}
