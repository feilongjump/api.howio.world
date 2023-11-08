package bootstrap

import (
	"github.com/feilongjump/api.howio.world/app/http/middlewares"
	"net/http"
	"strings"

	"github.com/feilongjump/api.howio.world/routes"
	"github.com/gin-gonic/gin"
)

// SetupRoute 初始化路由
func SetupRoute(r *gin.Engine) {

	registerGlobalMiddleWare(r)

	routes.RegisterAPIRoutes(r)

	setup404Handler(r)
}

// registerGlobalMiddleWare 注册中间件
func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middlewares.Cors(),
	)
}

// setup404Handler 404 路由
func setup404Handler(r *gin.Engine) {
	// 处理 404 请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		// 用以区分使用什么格式进行返回错误信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
