package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	DBUrl             string `mapstructure:"DB_URL"`
	ProductServiceUrl string `mapstructure:"SVC_URL_PRODUCT"`

	JWTSecretKey       string `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationHours int64  `mapstructure:"JWT_EXPIRATION_HOURS"`

	SMTP_FROM string `mapstructure:"SMTP_FROM"`
	SMTP_HOST string `mapstructure:"SMTP_HOST"`
	SMTP_PORT string `mapstructure:"SMTP_PORT"`
	SMTP_USER string `mapstructure:"SMTP_USER"`
	SMTP_PASS string `mapstructure:"SMTP_PASS"`
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
