package config

import "github.com/spf13/viper"

var (
	MySQLDSN      string
	RedisAddr     string
	RedisPassword string
	MinIOEndpoint string
	MinIOAccess   string
	MinIOSecret   string
	MinIOBucket   string
	JWTSecret     string
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	MySQLDSN = viper.GetString("mysql.dsn")
	RedisAddr = viper.GetString("redis.addr")
	RedisPassword = viper.GetString("redis.password")
	MinIOEndpoint = viper.GetString("minio.endpoint")
	MinIOAccess = viper.GetString("minio.access_key")
	MinIOSecret = viper.GetString("minio.secret_key")
	MinIOBucket = viper.GetString("minio.bucket")
	JWTSecret = viper.GetString("jwt.secret")
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
