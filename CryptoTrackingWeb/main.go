package main

import (
	"CryptoTrackingSql/sqlc"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"time"
)

type BUSDPercent struct {
	ID        int32
	Time      time.Time
	Symbol    string
	Price     float64
	PrevPrice float64
	Percent   float64
}

func main() {
	err := loadConfig(".")
	if err != nil {
		log.Fatalf(err.Error())
	}
	GStrConn = fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)
	busdPercent := BUSDPercent{-1, time.Now(), "error", 0, 0, 0}
	templates := template.Must(template.ParseFiles("templates/WelcomeTemplate.html"))
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1.Open DB
		conn, err := sql.Open(GConfig.DBDriver, GStrConn)
		if err != nil {
			log.Println(err.Error())
			conn.Close()
			return
		}
		query := sqlc.New(conn)
		oneMinData, err := query.GetAll1MinPercent(context.Background())
		if err != nil {
			log.Println(err.Error())
			if err := templates.ExecuteTemplate(w, "WelcomeTemplate.html", busdPercent); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
			}
		} else {
			var arrBusdPercent []BUSDPercent
			for _, v := range oneMinData {
				var it BUSDPercent
				it.ID = v.ID
				it.Time = v.Time.Time
				it.Symbol = v.Symbol.String
				it.Price = v.Price.Float64
				it.PrevPrice = v.Prevprice.Float64
				it.Percent = v.Percent.Float64
				arrBusdPercent = append(arrBusdPercent, it)
			}
			if err := templates.ExecuteTemplate(w, "WelcomeTemplate.html", arrBusdPercent); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
			}
		}
		conn.Close()
	})

	// Start web server, và sét cổng 8080
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
