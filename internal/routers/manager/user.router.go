package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (*UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router

	// private router
	userRouterPrivate := Router.Group("/admin/user")

	{
		userRouterPrivate.POST("/active_user")
	}
}
