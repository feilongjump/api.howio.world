package user

import (
	"github.com/feilongjump/api.howio.world/internal/hash"
	"gorm.io/gorm"
)

// BeforeSave 更新模型前调用
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(user.Password) {
		user.Password = hash.BcryptHash(user.Password)
	}

	return
}
