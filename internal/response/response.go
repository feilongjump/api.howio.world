package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidatorUnprocessableEntity 参数验证错误的响应
func ValidatorUnprocessableEntity(ctx *gin.Context, errors map[string]string) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"errors": errors,
	})
}

// Unauthorized 未授权的响应
func Unauthorized(ctx *gin.Context, message ...string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"msg": defaultMessage("unauthorized", message...),
	})
}

// defaultMessage 默认信息
func defaultMessage(defaultMessage string, message ...string) string {
	if len(message) > 0 {
		return message[0]
	}

	return defaultMessage
}
