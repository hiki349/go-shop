package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateRefreshToken(secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func CreateAccessToken(userID uuid.UUID, refreshToken, secret string) (string, error) {
	if err := VerifyToken(refreshToken, secret); err != nil {
		return "", fmt.Errorf("invalid refresh token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"id":  userID,
			"exp": time.Now().Add(time.Minute * 15).Unix(),
		})

	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func VerifyToken(tokenString string, secret string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
