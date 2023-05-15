package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MySQLDatabase          string `mapstructure:"MYSQLDATABASE"`
	MySQLHost              string `mapstructure:"MYSQLHOST"`
	MySQLUser              string `mapstructure:"MYSQLUSER"`
	MySQLPassword          string `mapstructure:"MYSQLPASSWORD"`
	MySQLPort              string `mapstructure:"MYSQLPORT"`
	Port                   string `mapstructure:"PORT"`
	AccessTokenPrivateKey  string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
