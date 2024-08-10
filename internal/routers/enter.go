package routers

import (
	"github.com/marktran77/go-ecomerce-backend-api/internal/routers/manager"
	"github.com/marktran77/go-ecomerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
