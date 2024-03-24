package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateRefreshToken(userID uuid.UUID, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
	token.Header["id"] = userID
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func CreateAccessToken(refreshToken *jwt.Token, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"exp": time.Now().Add(time.Minute * 15).Unix(),
		})
	token.Header["id"] = refreshToken.Header["id"]
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func VerifyToken(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
