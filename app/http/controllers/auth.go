package controllers

import (
	"github.com/feilongjump/api.howio.world/app/http/requests"
	userModel "github.com/feilongjump/api.howio.world/app/models/user"
	"github.com/feilongjump/api.howio.world/internal/jwt"
	"github.com/feilongjump/api.howio.world/internal/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (*AuthController) SignIn(ctx *gin.Context) {

	params := requests.SignInRequest{}
	if ok := requests.Validator(ctx, &params, params.ErrorMessage()); !ok {
		return
	}

	user, err := userModel.GetByUsername(params.Username)
	if err != nil {
		response.Unauthorized(ctx, "账号不存在")
		return
	}
	if !user.ComparePassword(params.Password) {
		response.Unauthorized(ctx, "密码错误")
		return
	}

	authToken(user, ctx)
}

func (*AuthController) SignUp(ctx *gin.Context) {

	params := requests.SignUpRequest{}
	if ok := requests.Validator(ctx, &params, params.ErrorMessage()); !ok {
		return
	}

	user := userModel.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	}
	if err := user.Create(); err != nil {
		response.Unauthorized(ctx, "注册用户失败")
		return
	}

	authToken(user, ctx)
}

func authToken(user userModel.User, ctx *gin.Context) {

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		response.Unauthorized(ctx)
		return
	}

	response.Success(ctx, gin.H{
		"token":      token,
		"token_type": "Bearer",
	})
}
