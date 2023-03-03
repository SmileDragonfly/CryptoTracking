package main

import "github.com/spf13/viper"

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	HostName string `mapstructure:"HOST_NAME"`
	HostPort int    `mapstructure:"HOST_PORT"`
	UserName string `mapstructure:"USER_NAME"`
	Password string `mapstructure:"PASS_WORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

var GConfig Config

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
