package user

import "github.com/feilongjump/api.howio.world/internal/database"

func GetByUsername(username string) (user User, err error) {

	if err = database.DB.Where("name = ? or email = ?", username, username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (user *User) Create() error {
	if err := database.DB.Create(&user); err != nil {
		return err.Error
	}

	return nil
}
