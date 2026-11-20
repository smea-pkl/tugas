package auth_middleware

import "github.com/golang-jwt/jwt/v5"

var secretKey = []byte("rahasia")

type CustomClaims struct {
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(cc *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	return token.SignedString(secretKey)
}
