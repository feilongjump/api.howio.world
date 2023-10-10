package jwt

import (
	"errors"
	"time"

	"github.com/feilongjump/api.howio.world/internal/config"
	"github.com/feilongjump/api.howio.world/internal/utils"
	jwtpkg "github.com/golang-jwt/jwt/v5"
)

// JWTCustomClaims 自定义载荷
type JWTCustomClaims struct {
	UserID uint64 `json:"user_id"`

	jwtpkg.RegisteredClaims
}

// GenerateToken 生成 jwt 令牌
func GenerateToken(UserID uint64) (string, error) {

	appName := config.GetString("app.name")
	expiresAt := time.Now().Add(time.Second * time.Duration(config.GetInt("jwt.ttl")))

	claims := JWTCustomClaims{
		UserID,
		jwtpkg.RegisteredClaims{
			Issuer:    appName,                                            // 签发者
			Subject:   "API Token",                                        // 签发主题
			Audience:  jwtpkg.ClaimStrings{appName + "_APP"},              // 签发受众
			ExpiresAt: jwtpkg.NewNumericDate(expiresAt),                   // 过期时间
			NotBefore: jwtpkg.NewNumericDate(time.Now().Add(time.Second)), // 最早使用时间
			IssuedAt:  jwtpkg.NewNumericDate(time.Now()),                  // 签发时间
			ID:        utils.GenerateRandomString(12),                     // wt ID, 类似于盐值
		},
	}

	// 使用特定的加密方式进行加密
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)

	// 使用指定的 secret 进行签名加密
	return token.SignedString([]byte(config.GetString("jwt.secret")))
}

// ParseToken 解析 jwt 令牌
func ParseToken(tokenString string) (*JWTCustomClaims, error) {

	token, err := jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(t *jwtpkg.Token) (interface{}, error) {
		return []byte(config.GetString("jwt.secret")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
