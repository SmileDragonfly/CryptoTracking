package main

import (
	"cryptoapi/logger"
	"cryptosql/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	// FIX-ERROR: Access to XMLHttpRequest at 'http://localhost:8888/5minup' from origin 'http://localhost:63342'
	// has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present on the requested resource
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll1MinPercentDesc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get5MinUp(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll5MinPercentDesc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get10MinUp(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll10MinPercentDesc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get15MinUp(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll15MinPercentDesc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get30MinUp(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll30MinPercentDesc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get60MinUp(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll60MinPercentDesc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get1MinDown(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll1MinPercentAsc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get5MinDown(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll5MinPercentAsc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get10MinDown(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll10MinPercentAsc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get15MinDown(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll15MinPercentAsc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get30MinDown(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll30MinPercentAsc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}

func get60MinDown(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 1.Open DB
	conn, err := sql.Open(GConfig.DBDriver, GStrConn)
	if err != nil {
		logger.Logger.Error(err.Error())
		conn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
		return
	}
	defer conn.Close()
	query := sqlc.New(conn)
	dataDB, err := query.GetAll60MinPercentAsc(c, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
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
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Marshal data to json failed"})
		} else {
			c.JSON(http.StatusOK, dataRet)
		}
	}
}
