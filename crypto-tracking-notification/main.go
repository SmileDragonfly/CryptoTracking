package main

import (
	"bytes"
	"context"
	"cryptonoti/logger"
	"cryptosql/sqlc"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

func main() {
	// Init logger
	err := logger.NewLogger("./config/logcfg.json")
	if err != nil {
		panic(err)
	}
	// Load config
	_, err = LoadConfig("./config")
	if err != nil {
		logger.Instance.Info("Load config error:", err)
		panic(err)
	}
	// Load threshold
	_, err = LoadThreshold("./config/threshold.json")
	if err != nil {
		logger.Instance.Info("Load threshold error:", err)
		panic(err)
	}
	// Create ticker to check DB interval
	ticker := time.NewTicker(time.Duration(config.ScanDBInterval) * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		var arrTopCoin3Min [][]TopCoin1Min
		for {
			// A select blocks until one of its cases can run, then it executes that case.
			// It chooses one at random if multiple are ready
			select {
			case <-done:
				return
			case <-ticker.C:
				func() {
					logger.Instance.Info("Start scan DB to send notification")
					db, err := sql.Open(config.DBDriver, config.GetConnectionString())
					if err != nil {
						logger.Instance.Error("Open DB connection failed:", err)
						return
					}
					defer db.Close()
					query := sqlc.New(db)
					upCoin1Min, err := query.GetAll1MinPercentDesc(context.Background(), 5)
					if err != nil {
						logger.Instance.Error("Get up coin 1 min failed:", err)
						return
					}
					var arrTopCoin1Min []TopCoin1Min
					for _, v := range upCoin1Min {
						if v.Percent.Float64 > threshold.OneMinUp {
							var it TopCoin1Min
							it.Symbol = v.Symbol.String
							it.Percent = v.Percent.Float64
							arrTopCoin1Min = append(arrTopCoin1Min, it)
						}
					}
					if len(arrTopCoin1Min) > 0 {
						arrTopCoin3Min = append(arrTopCoin3Min, arrTopCoin1Min)
						//byteTopCoin3Min, _ := json.MarshalIndent(&arrTopCoin3Min, "", "  ")
						byteTopCoin3Min, _ := json.Marshal(&arrTopCoin3Min)
						logger.Instance.Info("Data len: ", len(arrTopCoin3Min))
						logger.Instance.Info(string(byteTopCoin3Min))
						if len(arrTopCoin3Min) == 3 {
							// Append 3 min to a slice
							var arrTopCoinAll []TopCoin1Min
							for _, v := range arrTopCoin3Min {
								arrTopCoinAll = append(arrTopCoinAll, v...)
							}
							// Find if appear 3 times (3 minutes in a rows)
							mapCoinCount := make(map[string]int)
							mapCointPercent := make(map[string]float64)
							for _, v := range arrTopCoinAll {
								mapCoinCount[v.Symbol]++
								mapCointPercent[v.Symbol] += v.Percent
								if mapCoinCount[v.Symbol] == 3 {
									it := TopCoin1Min{
										Symbol:  v.Symbol,
										Percent: mapCointPercent[v.Symbol],
									}
									byteSend, _ := json.Marshal(&it)
									SendTeleMessage(config.TelegramUserId, string(byteSend))
								}
							}
							logger.Instance.Info("3 mins have no top coin, delete first item")
							// Delete first item in slice
							arrTopCoin3Min = append(arrTopCoin3Min[:0], arrTopCoin3Min[1:]...)
						}
					} else {
						logger.Instance.Info("1 min have no top coin, delete buffer")
						arrTopCoin3Min = nil
					}

				}()
			}
		}
	}()
	<-done
}

type TopCoin1Min struct {
	Symbol  string
	Percent float64
}

func SendTeleMessage(chatId int64, text string) error {
	link := "https://api.telegram.org/bot5466150074:AAEIrdGxJYOuGIfP6F2I5LN4DEE7mBql7Rc/sendMessage"
	data := struct {
		ChatId int64  `json:"chat_id"`
		Text   string `json:"text"`
	}{
		ChatId: chatId,
		Text:   text,
	}
	reqByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	logger.Instance.Info("Request send message:", string(reqByte))
	resp, err := http.Post(link, "application/json", bytes.NewBuffer(reqByte))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		strErr := fmt.Sprintf("Send failed. Status: %q", resp.StatusCode)
		logger.Instance.Error(strErr)
		return errors.New(strErr)
	}
	logger.Instance.Info("Send message successfully")
	return nil
}
