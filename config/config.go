package config

import (
	"github.com/spf13/viper"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Name     string
	Password string
}

type CurrencyApi struct {
	ApiKey        string
	UrlLatesRates string
	IntervalCall  int
	TimeoutCall   int
}

type ConfigSetting struct {
	Database    Database
	CurrencyApi CurrencyApi
}

var Setting ConfigSetting

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./../../config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err := viper.Unmarshal((&Setting)); err != nil {
		panic(err)
	}
}
