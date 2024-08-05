package initialize

import (
	"github.com/marktran77/go-ecomerce-backend-api/global"
	"github.com/marktran77/go-ecomerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
