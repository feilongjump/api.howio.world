package controllers

import (
	"net/http"

	userModel "github.com/feilongjump/api.howio.world/app/models/user"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (*UserController) Me(ctx *gin.Context) {

	user := userModel.Get(ctx.MustGet("user_id").(uint64))
	ctx.JSON(http.StatusOK, user)
}
