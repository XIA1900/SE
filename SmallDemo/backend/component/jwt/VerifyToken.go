package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(tokenContent string) error {
	if len(tokenContent) == 0 {
		return errors.New("no token error")
	}

	claims := &Claims{}

	verification, err := jwt.ParseWithClaims(tokenContent, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.New("signature error")
		}
		return errors.New(err.Error())
	}
	if !verification.Valid {
		return errors.New("verifying failure")
	}

	return nil
}
