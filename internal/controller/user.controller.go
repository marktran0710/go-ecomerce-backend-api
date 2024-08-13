package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marktran77/go-ecomerce-backend-api/internal/services"
	"github.com/marktran77/go-ecomerce-backend-api/pkg/response"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"hautran"})
}
