package token

import (
	"api_gateway/config"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaims(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Load().SIGNING_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
