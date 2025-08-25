package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/hyouhyan/gin-jwt-sample/auth"
	"github.com/hyouhyan/gin-jwt-sample/model"
)

func LoginHandler(c *gin.Context) {

	var inputUser model.User

	// リクエストからユーザー情報を取得
	if err := c.BindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユーザー情報の検証
	if inputUser.Username != model.ValidUser.Username || inputUser.Password != model.ValidUser.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}

	// トークンの発行（ヘッダー・ペイロード）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": inputUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(auth.SECRET_KEY))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error signing token"})
		return
	}

	// ヘッダーにトークンをセット
	c.Header("Authorization", tokenString)
	c.JSON(http.StatusOK, gin.H{"message": "login success"})
}
