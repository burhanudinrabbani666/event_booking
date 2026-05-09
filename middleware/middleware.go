package middleware

import (
	"event_booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"status": "NOT AUTHORIZED!",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"status": "NOT AUTHORIZED! TOKEN NOT VALID.",
		})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()

}
