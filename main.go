package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyouhyan/gin-jwt-sample/auth"
	"github.com/hyouhyan/gin-jwt-sample/handler"
)

func main() {
	r := gin.Default()

	r.POST("/login", handler.LoginHandler)

	authGroup := r.Group("/auth")
	authGroup.Use(auth.AuthMiddleware)
	authGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "you are authorized"})
	})

	r.Run(":8080")
}
