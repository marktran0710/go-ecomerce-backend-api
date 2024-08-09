package initialize

import (
	"context"
	"fmt"

	"github.com/marktran77/go-ecomerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r:= global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Host,
		Password: r.Password, // no password set
		DB:       r.Database,  // use default DB
		PoolSize: 10
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization Error:", zap.Error(err))
	}

	fmt.Println("Init Redis is running")
}
