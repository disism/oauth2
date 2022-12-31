package oauth2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	api := gin.Default()
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api.Use(CORS())

	oauth := api.Group("/oauth")
	oauth.GET("/token", GetTokenHandler)
	oauth.POST("/signup", SignupHandler)
	oauth.POST("/auth", AuthHandler)

	v1 := api.Group("/api/v1")
	v1.Use(Auth)
	v1.GET("/authorize", AuthorizeHandler)
	v1.POST("/applications/new", NewApplicationHandler)

	return api
}
