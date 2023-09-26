package requests

type SignInRequest struct {
	Username string `json:"username" binding:"required,gte=2"`
	Password string `json:"password" binding:"required,gte=6"`
}
