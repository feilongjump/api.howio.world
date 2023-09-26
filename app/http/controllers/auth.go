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
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
