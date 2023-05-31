package main

import "github.com/spf13/viper"

type Config struct {
	DBDriver              string `mapstructure:"DB_DRIVER"`
	HostName              string `mapstructure:"DB_HOST_NAME"`
	HostPort              int    `mapstructure:"DB_HOST_PORT"`
	UserName              string `mapstructure:"DB_USER_NAME"`
	Password              string `mapstructure:"DB_PASS_WORD"`
	DBName                string `mapstructure:"DB_NAME"`
	TickerGetPrice        int    `mapstructure:"TICKER_GET_PRICE"`
	NumberOfTopCoinRecord int    `mapstructure:"NUMBER_OF_TOPCOIN_RECORD"`
}

var GConfig Config
var GStrConn string

func loadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("Config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&GConfig)
	return err
}
