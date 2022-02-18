package auth

import (
	"GFBackend/config"
	"github.com/golang-jwt/jwt"
	"time"
)

type Payload struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type NewCookieInfo struct {
	Token   string
	Expires time.Time
}

func TokenGenerate(username string) (NewCookieInfo, error) {
	expirationTime := time.Now().Add(time.Duration(config.AppConfig.JWT.Expires) * time.Minute)
	payload := &Payload{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenContent, err := token.SignedString([]byte(config.AppConfig.JWT.SecretKey))
	if err != nil {
		return NewCookieInfo{}, err
	}
	newCookieInfo := NewCookieInfo{
		Token:   tokenContent,
		Expires: expirationTime,
	}
	return newCookieInfo, nil
}

func TokenRefresh(tokenContent string) (NewCookieInfo, error, bool) {
	payload := &Payload{}
	_, err := jwt.ParseWithClaims(tokenContent, payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWT.SecretKey), nil
	})

	if time.Unix(payload.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return NewCookieInfo{}, nil, false
	}

	newCookieInfo, err := TokenGenerate(payload.Username)
	if err != nil {
		return NewCookieInfo{}, err, false
	}

	return newCookieInfo, nil, true
}

func TokenVerify(token string) bool {
	payload := &Payload{}
	verification, err := jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWT.SecretKey), nil
	})
	if err != nil || !verification.Valid {
		return false
	}
	return true
}
