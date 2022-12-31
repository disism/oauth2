package oauth2

import (
	"github.com/disism/oauth2/config"
	"github.com/disism/oauth2/jwt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Auth(c *gin.Context) {
	a := c.Request.Header.Get("Authorization")
	if a == "" {
		c.JSON(500, gin.H{
			"status": "NOT CARRYING TOKEN",
		})
		c.Abort()
		return
	}
	_, err := ParseAuthorization(a)
	if err != nil {
		c.JSON(401, gin.H{})
		c.Abort()
		return
	}

	c.Next()
}

func ParseAuthorization(a string) (*jwt.Claims, error) {
	var (
		token  = strings.Split(a, "Bearer ")[1]
		secret = config.GetAuthTokenSecret()
	)
	parse, err := jwt.NewParseJWTToken(token, secret).JWTTokenParse()
	if err != nil {
		return nil, err
	}
	return parse, nil
}

func GetUserId(c *gin.Context) uint {
	a := c.Request.Header.Get("Authorization")
	if a == "" {
		c.JSON(500, gin.H{})
		return 0
	}
	parse, err := ParseAuthorization(a)
	if err != nil {
		c.JSON(401, gin.H{})
		return 0
	}

	atoi, err := strconv.Atoi(parse.Id)
	if err != nil {
		return 0
	}
	return uint(atoi)
}

// CORS For middleware of gin network framework, solve cross-domain problems.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
