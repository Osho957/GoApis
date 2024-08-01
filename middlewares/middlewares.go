package middlewares

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"osho.com/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token required"})
		return
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
