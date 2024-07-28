package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/marktran77/go-ecomerce-backend-api/internal/controller"
	"github.com/marktran77/go-ecomerce-backend-api/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func DD() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> DD")
		c.Next()
		fmt.Println("After --> DD")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")
}

func NewRouter() *gin.Engine {

	r := gin.Default()

	// use middlewares
	r.Use(middlewares.AuthMiddlewares(), AA(), BB(), CC)

	v1 := r.Group("v1/group")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/id", c.NewUserController().GetUserByID)
		// v1.PATCH("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.HEAD("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
	}

	// v2 := r.Group("v2/group")
	// {
	// 	v2.GET("/ping", Pong)
	// 	v2.PUT("/ping", Pong)
	// 	v2.PATCH("/ping", Pong)
	// 	v2.DELETE("/ping", Pong)
	// 	v2.HEAD("/ping", Pong)
	// 	v2.OPTIONS("/ping", Pong)
	// }

	return r
}
