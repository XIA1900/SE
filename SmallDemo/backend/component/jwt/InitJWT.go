package jwt

import (
	"backend/config"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey []byte
var expires int

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func InitJWT() {
	appConfig := config.GetAppConfig()
	secretKey = []byte(appConfig.JWT.SecretKey)
	expires = appConfig.JWT.Expires
}
