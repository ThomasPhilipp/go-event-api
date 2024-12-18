package middlewares

import (
	"net/http"

	"github.com/ThomasPhilipp/go-event-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	// verify token
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		//context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		//context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userId", userId) // add properties for later use

	context.Next() // call next handler
}
