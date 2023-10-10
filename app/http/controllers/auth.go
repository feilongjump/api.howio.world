package controllers

import (
	"net/http"

	"github.com/feilongjump/api.howio.world/app/http/requests"
	userModel "github.com/feilongjump/api.howio.world/app/models/user"
	"github.com/feilongjump/api.howio.world/internal/jwt"
	"github.com/feilongjump/api.howio.world/internal/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (*AuthController) SignIn(ctx *gin.Context) {

	params := requests.SignInRequest{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ValidatorUnprocessableEntity(ctx, params.GetErrors(err))
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

	authToken(user, ctx)
}

func (*AuthController) SignUp(ctx *gin.Context) {

	params := requests.SignUpRequest{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ValidatorUnprocessableEntity(ctx, params.GetErrors(err))
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

	authToken(user, ctx)
}

func authToken(user userModel.User, ctx *gin.Context) {

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		response.Unauthorized(ctx)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":      token,
		"token_type": "Bearer",
	})
}
