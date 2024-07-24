package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marktran77/go-ecomerce-backend-api/internal/services"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		"users":   []string{"cr7", "m10", "hautran"},
	})
}
