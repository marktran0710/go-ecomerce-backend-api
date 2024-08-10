package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (*AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := Router.Group("/admin/user")

	{
		adminRouterPrivate.POST("/active_user")
	}
}
