package authentication

import (
	"api/src/config"
	"time"

	"github.com/golang-jwt/jwt"
)

// CriateToken Criate JWT token
func CriateToken(userID uint64) (string, error) {

	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)
	return token.SignedString([]byte(config.SecretKey))
}
