package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jivelocity/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorization"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorization"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
