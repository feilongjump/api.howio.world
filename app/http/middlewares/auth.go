package middlewares

import (
	"strings"
	"time"

	"github.com/feilongjump/api.howio.world/internal/jwt"
	"github.com/feilongjump/api.howio.world/internal/response"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			response.Unauthorized(ctx)
			return
		}

		tokenSplit := strings.Fields(token)
		// 令牌类型是否正确
		if tokenSplit[0] != "Bearer" {
			response.Unauthorized(ctx)
			return
		}

		// 令牌可否正常解析
		claims, err := jwt.ParseToken(tokenSplit[1])
		if err != nil {
			response.Unauthorized(ctx)
		}

		// 令牌是否已过期
		if claims.ExpiresAt.Unix() <= time.Now().Unix() {
			response.Unauthorized(ctx)
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
