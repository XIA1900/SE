package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
)

var SecretKey = []byte("RoadCenter")

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func TestServer() {
	// controller registration
	http.HandleFunc("/signin", SignInService)
	http.HandleFunc("/welcome", WelcomeService)
	http.HandleFunc("/refresh", RefreshService)

	// start
	log.Fatal(http.ListenAndServe(":10012", nil))
}
