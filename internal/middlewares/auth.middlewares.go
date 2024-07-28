package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/marktran77/go-ecomerce-backend-api/pkg/response"
)

func AuthMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrCodeParamInvalid, "")
			c.Abort()
			return
		}

		c.Next()
	}
}
