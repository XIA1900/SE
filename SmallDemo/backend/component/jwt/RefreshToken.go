package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func RefreshToken(tokenContent string) (string, time.Time, error) {
	errSign := VerifyToken(tokenContent)
	if errSign != nil {
		return "", time.Now(), errors.New(errSign.Error())
	}

	claims := &Claims{}

	if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) > 30*time.Second {
		return "", time.Now(), errors.New("no need to refresh")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenContent, err := token.SignedString(secretKey)
	if err != nil {
		return "", time.Now(), err
	}

	return newTokenContent, expirationTime, nil
}
