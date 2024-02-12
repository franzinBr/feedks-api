package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) CreateToken(secret string, expiry int, claims map[string]any) (tokenSigned string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	jwtClaims := jwt.MapClaims{
		"exp": exp,
	}

	for key, value := range claims {
		jwtClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	tokenSigned, err = token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}
	return tokenSigned, err
}

func (s *TokenService) VerifyToken(tokenStr string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method:: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

}

func (s *TokenService) GetClaims(tokenStr string, secret string) (map[string]any, error) {
	token, err := s.VerifyToken(tokenStr, secret)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}

	claimMap := map[string]any{}

	for k, v := range claims {
		claimMap[k] = v
	}

	return claimMap, nil

}
