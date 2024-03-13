package main

import (
	"fmt"

	"github.com/feilongjump/api.howio.world/bootstrap"
	"github.com/feilongjump/api.howio.world/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置文件信息
	bootstrap.SetupConfig()
	// 初始化数据库连接
	bootstrap.SetupDatabase()
	// 初始化 redis 连接
	bootstrap.SetupRedisDB()

	// 初始化 Gin 实例
	router := gin.New()
	// 初始化路由
	bootstrap.SetupRoute(router)

	err := router.SetTrustedProxies(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 运行服务
	addr := fmt.Sprintf("%s:%s", config.GetString("app.host"), config.GetString("app.port"))
	if config.GetString("cert_file") != "" && config.GetString("key_file") != "" {
		err = router.RunTLS(addr, config.GetString("cert"), config.GetString("cert"))
	} else {
		err = router.Run(addr)
	}
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
