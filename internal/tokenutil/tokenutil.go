package tokenutil

import (
	"fmt"
	"log"
	"time"

	"github.com/gabrielfmcoelho/platform-core/domain"
	jwt "github.com/golang-jwt/jwt/v4"
)

// parse gorm ID of type uint to hex string
func parseUintToHex(id uint) string {
	return fmt.Sprintf("%x", id)
}

func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * time.Duration(expiry))
	claims := &domain.JwtCustomClaims{
		UserRoleID: user.Role.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        parseUintToHex(user.ID),
			Subject:   user.Bio.FirstName,
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * time.Duration(expiry))
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		UserRoleID: user.Role.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        parseUintToHex(user.ID),
			Subject:   user.Bio.FirstName,
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	log.Default().Println("Extracting ID from token", token)
	claims, ok := token.Claims.(jwt.MapClaims)
	log.Default().Println("Extracting ID from token", claims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims["id"].(string), nil
}

func SkipTokenValidation(requestToken string) (bool, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(requestToken, jwt.MapClaims{})
	if err != nil {
		return false, err
	}
	// if token has claim "admin" and it is true, return the true
	if token.Claims.(jwt.MapClaims)["apiAdmin"] == true {
		return true, nil
	}
	return false, nil
}
