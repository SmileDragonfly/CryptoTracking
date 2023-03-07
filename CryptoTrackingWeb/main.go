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

type AllTimePercent struct {
	OneMin     []BUSDPercent
	FiveMin    []BUSDPercent
	TenMin     []BUSDPercent
	FifteenMin []BUSDPercent
	ThirtyMin  []BUSDPercent
	SixtyMin   []BUSDPercent
}

func main() {
	err := loadConfig(".")
	if err != nil {
		log.Fatalf(err.Error())
	}
	GStrConn = fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		GConfig.HostPort, GConfig.HostName, GConfig.UserName, GConfig.Password, GConfig.DBName)
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
		var allTimePercent AllTimePercent
		query := sqlc.New(conn)
		oneMinData, err := query.GetAll1MinPercentDesc(context.Background(), 10)
		if err != nil {
			log.Println(err.Error())
			if err := templates.ExecuteTemplate(w, "WelcomeTemplate.html", allTimePercent); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
			}
		} else {
			for _, v := range oneMinData {
				var it BUSDPercent
				it.ID = v.ID
				it.Time = v.Time.Time
				it.Symbol = v.Symbol.String
				it.Price = v.Price.Float64
				it.PrevPrice = v.Prevprice.Float64
				it.Percent = v.Percent.Float64
				allTimePercent.OneMin = append(allTimePercent.OneMin, it)
			}
		}

		FiveMinData, err := query.GetAll5MinPercentDesc(context.Background(), 10)
		if err != nil {
			log.Println(err.Error())
			if err := templates.ExecuteTemplate(w, "WelcomeTemplate.html", allTimePercent); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
			}
		} else {
			for _, v := range FiveMinData {
				var it BUSDPercent
				it.ID = v.ID
				it.Time = v.Time.Time
				it.Symbol = v.Symbol.String
				it.Price = v.Price.Float64
				it.PrevPrice = v.Prevprice.Float64
				it.Percent = v.Percent.Float64
				allTimePercent.FiveMin = append(allTimePercent.FiveMin, it)
			}
		}

		TenMinData, err := query.GetAll10MinPercentDesc(context.Background(), 10)
		if err != nil {
			log.Println(err.Error())
			if err := templates.ExecuteTemplate(w, "WelcomeTemplate.html", allTimePercent); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
			}
		} else {
			for _, v := range TenMinData {
				var it BUSDPercent
				it.ID = v.ID
				it.Time = v.Time.Time
				it.Symbol = v.Symbol.String
				it.Price = v.Price.Float64
				it.PrevPrice = v.Prevprice.Float64
				it.Percent = v.Percent.Float64
				allTimePercent.TenMin = append(allTimePercent.TenMin, it)
			}
		}

		if err := templates.ExecuteTemplate(w, "WelcomeTemplate.html", allTimePercent); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
		}
		conn.Close()
	})

	// Start web server, và sét cổng 8080
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
