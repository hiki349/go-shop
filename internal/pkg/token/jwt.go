package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenUserInfo struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

type JwtUserInfoClaims struct {
	jwt.RegisteredClaims
	User *TokenUserInfo `json:"user,omitempty"`
}

func NewToken(secret string, expirationTime int, userInfo *TokenUserInfo) (string, error) {
	claims := JwtUserInfoClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationTime) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: userInfo,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(secret)

	return signedToken, err
}

func VerifyToken(secret string, token string) (*TokenUserInfo, bool) {
	t, err := jwt.ParseWithClaims(token, &JwtUserInfoClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, false
	}

	expTime, err := t.Claims.GetExpirationTime()
	if err != nil {
		return nil, false
	}

	if !t.Valid && expTime.Before(time.Now()) {
		return nil, false
	}

	if userInfo, ok := t.Claims.(*JwtUserInfoClaims); ok {
		return userInfo.User, true
	}

	return nil, false
}
