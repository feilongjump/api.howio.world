package hash

import "golang.org/x/crypto/bcrypt"

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	// GenerateFromPassword 的第二个参数是 cost 值。官方建议值在 10 - 15 内，而默认值是 DefaultCost int = 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err.Error())
	}

	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// BcryptIsHashed 判断字符串是否是哈希过的数据
func BcryptIsHashed(str string) bool {
	// bcrypt 加密后的长度等于 60
	return len(str) == 60
}
