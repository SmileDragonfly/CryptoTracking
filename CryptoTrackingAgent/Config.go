package main

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	HostName      string `mapstructure:"HOST_NAME"`
	HostPort      int    `mapstructure:"HOST_PORT"`
	UserName      string `mapstructure:"USER_NAME"`
	Password      string `mapstructure:"PASS_WORD"`
	DBName        string `mapstructure:"DB_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

var GConfig Config

func loadConfig(path string) error {
	viper.
}
