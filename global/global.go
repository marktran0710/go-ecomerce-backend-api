package global

import (
	"github.com/marktran77/go-ecomerce-backend-api/pkg/logger"
	"github.com/marktran77/go-ecomerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb *redis.Client
	Mdb    *gorm.DB
)

/*
Config
Redis
Mysql
*/
