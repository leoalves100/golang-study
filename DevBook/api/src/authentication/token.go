package authentication

import (
	"api/src/config"
	"net/http"
	"strings"
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

// ValidToken Checks if the passed token is valid
func ValidToken(r *http.Request) error {

	// tokenString := extractToken(r)

	return nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, "")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}
