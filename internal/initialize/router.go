package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/marktran77/go-ecomerce-backend-api/global"
	"github.com/marktran77/go-ecomerce-backend-api/internal/routers"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()

	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	r.Use() // logging
	r.Use() //cross
	r.Use() // global limiters

	manageRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}

	return r
}
