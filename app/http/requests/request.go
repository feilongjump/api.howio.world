package requests

import (
	"github.com/feilongjump/api.howio.world/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Validator 验证请求参数
func Validator(ctx *gin.Context, params any, errorMessage map[string]string) bool {
	if err := ctx.ShouldBindJSON(params); err != nil {
		errMap := make(map[string]string)

		for _, v := range err.(validator.ValidationErrors) {

			if errMessage, ok := errorMessage[v.Field()+"."+v.Tag()]; ok {
				errMap[v.Field()] = errMessage
			} else {
				errMap[v.Field()] = v.Error()
			}
		}

		response.ValidatorUnprocessableEntity(ctx, errMap)
		return false
	}

	return true
}
