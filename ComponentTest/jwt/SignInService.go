package jwt

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func SignInService(w http.ResponseWriter, r *http.Request) {
	var credential Credential
	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		return
	}

	if credential.Password != "123456" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: credential.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenContent, err := token.SignedString(SecretKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenContent,
		Expires: expirationTime,
	})

}
