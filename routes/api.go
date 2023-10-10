package routes

import (
	"net/http"

	"github.com/feilongjump/api.howio.world/app/http/controllers"
	"github.com/feilongjump/api.howio.world/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

// 注册 API 路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {
		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello World!",
		})
	})

	registerAuthRoutes(r)

	registerUserRoutes(r)
}

// registerAuthRoutes Auth API
func registerAuthRoutes(r *gin.Engine) {
	authController := new(controllers.AuthController)

	authRoute := r.Group("/auth")
	authRoute.POST("sign-in", authController.SignIn)
	authRoute.POST("sign-up", authController.SignUp)
}

// registerUserRoutes User API
func registerUserRoutes(r *gin.Engine) {
	userController := new(controllers.UserController)

	r.GET("/me", middlewares.Auth(), userController.Me)
}
