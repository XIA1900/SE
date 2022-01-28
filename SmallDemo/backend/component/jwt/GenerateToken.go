package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(username string) (string, time.Time, error) {
	expiresTime := time.Now().Add(time.Duration(int64(expires)) * time.Minute)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenContent, err := token.SignedString(secretKey)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenContent, expiresTime, nil
}
