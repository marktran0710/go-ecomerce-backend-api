package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (*UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userRouterPublic := Router.Group("/user")

	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}

	// private router
	userRoutePrivate := Router.Group("/user")
	// userRoutePrivate.Use(Limiter())
	// userRoutePrivate.Use(Authen())
	// userRoutePrivate.Use(Permission())
	{
		userRoutePrivate.GET("/get_info")
	}
}
