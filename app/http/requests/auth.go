package requests

type SignInRequest struct {
	Username string `json:"username" binding:"required,gte=5"`
	Password string `json:"password" binding:"required,gte=6"`
}

type SignUpRequest struct {
	Name             string `json:"name" binding:"required,gte=5"`
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,gte=6"`
	PasswordConfirm  string `json:"password_confirm" binding:"required,eqfield=Password"`
	VerificationCode string `json:"verification_code" binding:"required,number,len=6"`
}

func (SignInRequest) ErrorMessage() map[string]string {
	return map[string]string{
		"Username.required": "请输入用户名",
		"Username.gte":      "用户名长度不得小于 5",
		"Password.required": "请输入密码",
		"Password.gte":      "密码长度不得小于 6",
	}
}

func (SignUpRequest) ErrorMessage() map[string]string {
	return map[string]string{
		"Name.required":             "请输入用户名",
		"Name.gte":                  "用户名长度不得小于 5",
		"Email.required":            "请输入邮箱",
		"Email.email":               "邮箱格式不正确",
		"Password.required":         "请输入密码",
		"Password.gte":              "密码长度不得小于 6",
		"PasswordConfirm.required":  "请输入确认密码",
		"PasswordConfirm.eqfield":   "确认密码与密码不一致",
		"VerificationCode.required": "请输入邮箱验证码",
		"VerificationCode.number":   "邮箱验证码格式错误",
		"VerificationCode.len":      "邮箱验证码长度不正确",
	}
}
