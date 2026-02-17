package config

import "github.com/spf13/viper"

func Init() error {
	viper.SetConfigName("dsn")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}
	return nil
}

type Config struct {
	MySQL struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"mysql"`
}

var Cfg Config

func GetMySQLDSN() string {
	return Cfg.MySQL.DSN
}
