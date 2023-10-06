package controllers

import (
	"net/http"

	"github.com/feilongjump/api.howio.world/app/http/requests"
	userModel "github.com/feilongjump/api.howio.world/app/models"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (*AuthController) SignIn(ctx *gin.Context) {

	params := requests.SignInRequest{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user, err := userModel.GetByUsername(params.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "账号不存在",
		})
		return
	}
	if !user.ComparePassword(params.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "密码错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (*AuthController) SignUp(ctx *gin.Context) {

	params := requests.SignUpRequest{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user := userModel.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	}
	if err := user.Create(); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "注册用户失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
