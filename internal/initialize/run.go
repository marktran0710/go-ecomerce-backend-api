package initialize

import (
	"fmt"

	"github.com/marktran77/go-ecomerce-backend-api/global"
)

func Run() {
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
