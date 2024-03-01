// jwt.go

package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key") // Change this to a strong secret key

// CustomClaims represents the custom claims for the JWT
type CustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken generates a new JWT token for the provided user ID and username
func GenerateToken(userID, username string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
			Issuer:    "your-app-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken validates the provided JWT token
func ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
