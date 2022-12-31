package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type JWTParse struct {
	Token  string
	Secret string
}

type v interface {
	JWTTokenParse() (*Claims, error)
}

func NewParseJWTToken(token string, secret string) *JWTParse {
	return &JWTParse{Token: token, Secret: secret}
}

func (p *JWTParse) JWTTokenParse() (*Claims, error) {
	parse, err := jwt.ParseWithClaims(p.Token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(p.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parse.Claims.(*Claims); ok && parse.Valid {
		return claims, nil
	}
	return nil, errors.New("errors.ErrTokenInvalid")
}
