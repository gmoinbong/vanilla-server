package storage

import (
	"context"
	"fmt"
	"log"
	"vanilla-server/internal/config"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitDB(cfg *config.Config) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
		Password: cfg.DBPassword,
		DB:       cfg.DBName,
	})

	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}
	fmt.Println(status)
	return rdb, nil
}
