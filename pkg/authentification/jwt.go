package authentification

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var RefreshSecret = []byte("my_refresh_secret")

func GenerateToken(secret, email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(RefreshSecret)
}

func ParseToken(secret, tokenString string) (string, string, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		role, _ := claims["role"].(string)
		return email, role, nil
	}
	return "", "", err
}
