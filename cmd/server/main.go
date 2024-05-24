package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", Pong)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping",
	})
}
