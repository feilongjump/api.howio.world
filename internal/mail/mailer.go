package mail

// VerificationCode 发送邮箱验证码
func VerificationCode(to string, code string) bool {
	body := "您的验证码为：" + code

	return NewMail(to, "邮箱验证码", body).SendEmail()
}
