package user

import (
	"github.com/feilongjump/api.howio.world/internal/database"
	"github.com/feilongjump/api.howio.world/internal/hash"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func GetByUsername(username string) (user User, err error) {

	if err = database.DB.Where("name = ? or email = ?", username, username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (user *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, user.Password)
}
