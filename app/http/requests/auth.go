package requests

type SignInRequest struct {
	Username string `json:"username" binding:"required,gte=2"`
	Password string `json:"password" binding:"required,gte=6"`
}

type SignUpRequest struct {
	Name            string `json:"name" binding:"required,gte=2"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,gte=6"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}
