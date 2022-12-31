package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type Userdata struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Mail     string `json:"mail,omitempty"`
}

func NewUserdata(id string, username string, mail string) *Userdata {
	return &Userdata{Id: id, Username: username, Mail: mail}
}

type JWT interface {
	JWTTokenGenerator(secret string) (token string, err error)
}

// JWTTokenGenerator ...
func (c *Claims) JWTTokenGenerator(secret string) (token string, err error) {
	g, err := jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return g, nil
}
