package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	MathRand "math/rand"
)

// GenerateRandomString 生成随机加密字符串
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

// CreateSixCaptcha 创建 6 位数的随机数
func CreateSixCaptcha() string {
	return fmt.Sprintf("%06v", MathRand.Intn(1000000))
}
