package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/jeanphorn/log4go"
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
	// Load config
	err := loadConfig(".")
	if err != nil {
		log.Error(err.Error())
	}
	GStrConn = fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)

	// Start api
	router := gin.Default()
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowHeaders = append(cfg.AllowHeaders, "ngrok-skip-browser-warning")
	router.Use(cors.New(cfg))
	router.GET("/1minup", get1MinUp)
	router.GET("/5minup", get5MinUp)
	router.GET("/10minup", get10MinUp)
	router.GET("/15minup", get15MinUp)
	router.GET("/30minup", get30MinUp)
	router.GET("/60minup", get60MinUp)
	router.GET("/1mindown", get1MinDown)
	router.GET("/5mindown", get5MinDown)
	router.GET("/10mindown", get10MinDown)
	router.GET("/15mindown", get15MinDown)
	router.GET("/30mindown", get30MinDown)
	router.GET("/60mindown", get60MinDown)
	router.Run("localhost:8888")
}
