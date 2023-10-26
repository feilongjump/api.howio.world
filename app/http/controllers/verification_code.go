package controllers

import (
	"github.com/feilongjump/api.howio.world/app/http/requests"
	"github.com/feilongjump/api.howio.world/internal/mail"
	"github.com/feilongjump/api.howio.world/internal/response"
	"github.com/feilongjump/api.howio.world/internal/utils"
	"github.com/gin-gonic/gin"
)

type VerificationCodeController struct{}

func (vc *VerificationCodeController) VerificationCode(ctx *gin.Context) {

	if ctx.Param("medium") == "email" {
		vc.sendMailVerificationCode(ctx)
		return
	}

	response.NotFound(ctx)
}

func (*VerificationCodeController) sendMailVerificationCode(ctx *gin.Context) {

	params := requests.SendMailVerificationCodeRequest{}
	if ok := requests.Validator(ctx, &params, params.ErrorMessage()); !ok {
		return
	}

	mail.VerificationCode(params.Email, utils.CreateSixCaptcha())

	response.SuccessNoContent(ctx)
}
