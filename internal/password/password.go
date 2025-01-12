package password

import (
	"github.com/gabrielfmcoelho/platform-core/domain"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the raw password
func HashPassword(rawPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.ErrInternalServerError
	}
	return string(hash), nil
}

// VerifyPassword verifies the raw password with the hashed password
func VerifyPassword(hashedPassword, rawPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	if err != nil {
		return domain.ErrUserPasswordNotMatch
	}
	return nil
}
