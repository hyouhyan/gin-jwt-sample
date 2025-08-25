package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hyouhyan/gin-jwt-sample/config"
)

func AuthMiddleware(c *gin.Context) {
	// Authorizationヘッダーからトークンを取得
	tokenString := c.GetHeader("Authorization")

	// トークンの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_KEY), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	c.Next()
}
