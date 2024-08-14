//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/marktran77/go-ecomerce-backend-api/internal/controller"
	"github.com/marktran77/go-ecomerce-backend-api/internal/repo"
	"github.com/marktran77/go-ecomerce-backend-api/internal/services"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		controller.NewUserController,
		services.NewUserService,
		repo.NewUserRepo,
	)

	return new(controller.UserController), nil
}
