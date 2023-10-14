package middlewares

import (
	"github.com/feilongjump/api.howio.world/internal/jwt"
	"github.com/feilongjump/api.howio.world/internal/response"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthUser struct {
	UserID uint64
}

// Auth 用户登录
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUser := AuthUser{}

		if ok := checkAuthToken(ctx, &authUser, true); !ok {
			return
		}

		ctx.Set("user_id", authUser.UserID)
	}
}

// Guest 游客登陆，验证是否存在用户 ID
func Guest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUser := AuthUser{}

		checkAuthToken(ctx, &authUser, false)

		ctx.Set("user_id", authUser.UserID)
	}
}

// CheckAuthToken 检查用户令牌
func checkAuthToken(ctx *gin.Context, authUser *AuthUser, isAbort bool) bool {

	token := ctx.GetHeader("Authorization")
	if token == "" {
		if isAbort {
			response.Unauthorized(ctx)
		}
		return false
	}

	tokenSplit := strings.Fields(token)
	// 令牌类型是否正确
	if tokenSplit[0] != "Bearer" {
		if isAbort {
			response.Unauthorized(ctx)
		}
		return false
	}

	// 令牌可否正常解析
	claims, err := jwt.ParseToken(tokenSplit[1])
	if err != nil {
		if isAbort {
			response.Unauthorized(ctx, err.Error())
		}
		return false
	}

	authUser.UserID = claims.UserID
	return true
}
