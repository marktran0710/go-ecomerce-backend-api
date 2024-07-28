package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()

}
