package routes

import (
	"net/http"

	"github.com/feilongjump/api.howio.world/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {
		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello World!",
		})
	})

	registerAuthRoutes(r)
}

// Auth
func registerAuthRoutes(r *gin.Engine) {
	authController := new(controllers.AuthController)

	authRoute := r.Group("/auth")
	authRoute.POST("sign-in", authController.SignIn)
	authRoute.POST("sign-up", authController.SignUp)
}
