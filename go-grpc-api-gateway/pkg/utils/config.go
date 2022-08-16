package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"SVC_URL_AUTH"`
	ProductSvcUrl string `mapstructure:"SVC_URL_PRODUCT"`
	OrderSvcUrl   string `mapstructure:"SVC_URL_ORDER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
