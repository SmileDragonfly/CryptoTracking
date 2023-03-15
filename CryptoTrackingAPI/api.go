package main

import (
	"CryptoTrackingSql/sqlc"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

func get1MinUp(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll1MinPercentDesc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get5MinUp(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll5MinPercentDesc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get10MinUp(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll10MinPercentDesc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get15MinUp(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll15MinPercentDesc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get30MinUp(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll30MinPercentDesc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get60MinUp(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll60MinPercentDesc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get1MinDown(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll1MinPercentAsc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get5MinDown(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll5MinPercentAsc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get10MinDown(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll10MinPercentAsc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get15MinDown(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll15MinPercentAsc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get30MinDown(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll30MinPercentAsc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}

func get60MinDown(c *gin.Context) {
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}
	query := sqlc.New(conn)
	dataDB, err := query.GetAll60MinPercentAsc(c, 10)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Get data from db failed"})
	} else {
		var dataRet []BUSDPercent
		for _, v := range dataDB {
			var it BUSDPercent
			it.ID = v.ID
			it.Time = v.Time.Time
			it.Symbol = v.Symbol.String
			it.Price = v.Price.Float64
			it.PrevPrice = v.Prevprice.Float64
			it.Percent = v.Percent.Float64
			dataRet = append(dataRet, it)
		}
		data, err := json.Marshal(dataRet)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			sData := string(data)
			c.JSON(http.StatusOK, sData)
		}
	}
}
