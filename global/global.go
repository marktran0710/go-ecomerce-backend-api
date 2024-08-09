package global

import (
	"github.com/marktran77/go-ecomerce-backend-api/pkg/logger"
	"github.com/marktran77/go-ecomerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)

/*
Config
Redis
Mysql
*/
