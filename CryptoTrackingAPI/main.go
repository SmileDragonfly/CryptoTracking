package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Init log file
	f, err := os.OpenFile("CryptoTrackingAPI.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("----------Start CryptoTrackingAPI----------")

	// Load config
	err = loadConfig(".")
	if err != nil {
		log.Fatalf(err.Error())
	}
	GStrConn = fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)

	// Start api
	router := gin.Default()
	router.GET("/1MinUp", get1MinUp)
	router.GET("/5MinUp", get5MinUp)
	router.GET("/10MinUp", get10MinUp)
	router.GET("/15MinUp", get15MinUp)
	router.GET("/30MinUp", get30MinUp)
	router.GET("/60MinUp", get60MinUp)
	router.GET("/1MinDown", get1MinDown)
	router.GET("/5MinDown", get5MinDown)
	router.GET("/10MinDown", get10MinDown)
	router.GET("/15MinDown", get15MinDown)
	router.GET("/30MinDown", get30MinDown)
	router.GET("/60MinDown", get60MinDown)
	router.Run("localhost:8888")
}
