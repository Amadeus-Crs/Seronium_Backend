package repository

import (
	"Seronium/internal/config"
	"Seronium/internal/model"
	"context"
	"fmt"

	"github.com/minio/minio-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	db, err := gorm.Open(mysql.Open(config.GetMySQLDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to mysql: %w", err)
	}
	DB = db

	return nil
}

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       0,
	})
}

var CTX = context.Background()
var MinIOClient *minio.Client

func InitMinIO() error {
	client, err := minio.New(config.MinIOEndpoint, config.MinIOAccess, config.MinIOSecret, false)
	if err != nil {
		return err
	}
	MinIOClient = client
	exists, err := MinIOClient.BucketExists(config.MinIOBucket)
	if err != nil {
		return err
	}
	if !exists {
		if err := MinIOClient.MakeBucket(config.MinIOBucket, ""); err != nil {
			return err
		}
	}
	return nil
}

func AutoMigrate() error {
	return DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}, &model.Like{}, &model.Collection{})
}
