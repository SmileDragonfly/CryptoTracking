package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBAddress      string `mapstructure:"DB_ADDRESS"`
	DBPort         int64  `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	ScanDBInterval int64  `mapstructure:"SCAN_DB_INTERVAL"`
	TelegramUserId int64  `mapstructure:"TELEGRAM_USER_ID"`
}

type Threshold struct {
	OneMinUp       float64 `json:"OneMinUp"`
	FiveMinUp      float64 `json:"FiveMinUp"`
	TenMinUp       float64 `json:"TenMinUp"`
	FifteenMinUp   float64 `json:"FifteenMinUp"`
	ThirtyMinUp    float64 `json:"ThirtyMinUp"`
	SixtyMinUp     float64 `json:"SixtyMinUp"`
	OneMinDown     float64 `json:"OneMinDown"`
	FiveMinDown    float64 `json:"FiveMinDown"`
	TenMinDown     float64 `json:"TenMinDown"`
	FifteenMinDown float64 `json:"FifteenMinDown"`
	ThirtyMinDown  float64 `json:"ThirtyMinDown"`
	SixtyMinDown   float64 `json:"SixtyMinDown"`
}

var config Config
var threshold Threshold

func LoadConfig(sPath string) (*Config, error) {
	viper.AddConfigPath(sPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// So we call viper.AutomaticEnv() to tell viper to automatically override values
	// that it has read from config file with the values of the corresponding environment variables
	// if they exist.
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	return &config, err
}

func (c Config) GetConnectionString() string {
	return fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBPort, c.DBAddress, c.DBUser, c.DBPassword, c.DBName)
}

func LoadThreshold(sPath string) (*Threshold, error) {
	data, err := os.ReadFile(sPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &threshold)
	if err != nil {
		return nil, err
	}
	return &threshold, nil
}
