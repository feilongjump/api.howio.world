package main

import (
	"fmt"

	"github.com/feilongjump/api.howio.world/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	router := gin.New()

	// 初始化路由
	bootstrap.SetupRoute(router)

	router.SetTrustedProxies(nil)
	// 运行服务
	err := router.Run(":3000")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
