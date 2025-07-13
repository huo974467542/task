package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func JWTAuth(c *gin.Context) {
	// 跳过不需要鉴权的路径
	if c.Request.URL.Path == "/register" {
		c.Next()
		return
	}

	tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("id", claims["id"])
		c.Set("username", claims["username"])
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	}
}
