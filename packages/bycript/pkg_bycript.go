package pkgbycript

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}

	return string(hashByte), nil
}

func CheckPasswordHash(password, HashingPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashingPassword), []byte(password))

	return err == nil
}
