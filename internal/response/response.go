package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success 成功
func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

// SuccessCreated 成功创建
func SuccessCreated(ctx *gin.Context, data any, message ...string) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": defaultMessage("success", message...),
		"data":    data,
	})
}

// SuccessNoContent 成功，但没有响应内容
func SuccessNoContent(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// NotFound 数据不存在
func NotFound(ctx *gin.Context, message ...string) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error": defaultMessage("The resource you want was not found", message...),
	})
}

// Unauthorized 未授权的响应
func Unauthorized(ctx *gin.Context, message ...string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": defaultMessage("Unauthorized", message...),
	})
}

// ValidatorUnprocessableEntity 参数验证错误的响应
func ValidatorUnprocessableEntity(ctx *gin.Context, errors map[string]string) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"error": errors,
	})
}

// InternalServerError 服务器错误
func InternalServerError(ctx *gin.Context, message ...string) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": defaultMessage("Server error", message...),
	})
}

// defaultMessage 默认信息
func defaultMessage(defaultMessage string, message ...string) string {
	if len(message) > 0 {
		return message[0]
	}

	return defaultMessage
}
