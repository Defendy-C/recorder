package auth

import "github.com/golang-jwt/jwt"

func NewJwtToken(payload map[string]interface{}, secretKey []byte) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = jwt.MapClaims(payload)

	return t.SignedString(secretKey)
}
