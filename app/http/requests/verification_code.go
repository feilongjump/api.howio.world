package requests

type SendMailVerificationCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (SendMailVerificationCodeRequest) ErrorMessage() map[string]string {
	return map[string]string{
		"Email.required": "请输入邮箱",
		"Email.email":    "邮箱格式有误",
	}
}
