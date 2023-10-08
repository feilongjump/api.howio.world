package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 未授权的响应
func Unauthorized(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"msg": "unauthorized",
	})
}
