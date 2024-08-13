package user

import (
	"github.com/gin-gonic/gin"
	"github.com/marktran77/go-ecomerce-backend-api/internal/controller"
	"github.com/marktran77/go-ecomerce-backend-api/internal/repo"
	"github.com/marktran77/go-ecomerce-backend-api/internal/services"
)

type UserRouter struct {
}

func (*UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	ur := repo.NewUserRepo()
	us := services.NewUserService(ur)
	userController := controller.NewUserController(us)

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		// userRouterPublic.POST("/otp")
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
