package user

import (
	"github.com/feilongjump/api.howio.world/app/models"
	"github.com/feilongjump/api.howio.world/internal/hash"
)

type User struct {
	models.BaseModel

	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

// ComparePassword 密码校验
func (user *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, user.Password)
}
