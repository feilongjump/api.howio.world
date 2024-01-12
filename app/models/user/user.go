package user

import (
	"github.com/feilongjump/api.howio.world/app/models"
	"github.com/feilongjump/api.howio.world/internal/hash"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(50);not null;index" json:"name"`
	Password string `gorm:"column:password;type:varchar(64);not null" json:"-"`
	Email    string `gorm:"column:email;type:varchar(100);not null;uniqueKey" json:"email"`

	models.BaseTimeModel
}

// ComparePassword 密码校验
func (user *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, user.Password)
}
