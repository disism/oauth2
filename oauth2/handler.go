package oauth2

import "C"
import (
	"fmt"
	"github.com/disism/oauth2/config"
	"github.com/disism/oauth2/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
)

// AuthorizeHandler When Oauth authentication is performed, the request will be made through this method.
func AuthorizeHandler(c *gin.Context) {
	responseType := c.Query("response_type")
	clientId := c.Query("client_id")
	redirectUri := c.Query("redirect_uri")
	//scope := c.Query("scope")

	code := uuid.New().String()
	if responseType != "code" {
		c.JSON(503, gin.H{
			"status": "ONLY LICENSE CODE MODE IS SUPPORTED FOR NOW",
		})
		return
	}

	// check is client_id?
	cid, err := strconv.Atoi(clientId)
	if err != nil {
		c.JSON(503, gin.H{
			"status": err.Error(),
		})
		return
	}

	client, err := NewClientId(uint(cid)).Get()
	if err != nil {
		c.JSON(503, gin.H{
			"status": err.Error(),
		})
		return
	}

	if redirectUri != client.AuthorizationCallbackURL {
		c.JSON(503, gin.H{
			"status": "THE CALLBACK ADDRESS MUST BE THE SAME AS WHEN THE CLIENT WAS REGISTERED, THIS IS A SECURITY CHECK",
		})
		return
	}

	if err := NewCache(1).SETCODE(code, strconv.Itoa(int(GetUserId(c)))); err != nil {
		c.JSON(503, gin.H{
			"status": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"callback": fmt.Sprintf("%s?code=%s", client.AuthorizationCallbackURL, code),
	})
}

// GetTokenHandler This method is used to obtain the token when Oauth authentication is performed .
func GetTokenHandler(c *gin.Context) {
	clientId := c.Query("client_id")
	clientSecret := c.Query("client_secret")
	grantType := c.Query("grant_type")
	code := c.Query("code")
	redirectUri := c.Query("redirect_uri")

	v := NewCache(1).GETUSERID(code)
	uid, err := strconv.Atoi(v)
	if err != nil {
		c.JSON(503, gin.H{
			"status": err.Error(),
		})
		return
	}
	get, err := NewOauthUsersId(uint(uid)).Get()
	if err != nil {
		c.JSON(503, gin.H{
			"status": err.Error(),
		})
		return
	}

	var (
		issuer = config.GetDomain()
		expir  = config.GetAuthTokenExpir()
	)

	fmt.Println(clientSecret)

	userdata := jwt.NewUserdata(v, get.Username, get.Mail)
	claims := jwt.NewRegisteredClaims(issuer, "", "", expir)
	generator, err := jwt.NewClaims(userdata, claims).JWTTokenGenerator(clientSecret)
	c.JSON(200, gin.H{
		"client_id":     clientId,
		"grant_type":    grantType,
		"access_token":  generator,
		"token_type":    "example",
		"expires_in":    3600,
		"refresh_token": "",
		"redirect_uri":  redirectUri,
	})
}

// SignupHandler Account registration by username, email, password.
func SignupHandler(c *gin.Context) {
	username := c.PostForm("username")
	mail := c.PostForm("mail")
	password := c.PostForm("password")

	if _, err := NewCreateOauthUsers(username, mail, password).Create(); err != nil {
		c.JSON(501, gin.H{
			"status": fmt.Sprintf("CREATE ACCOUNTS ERROR: %v", err.Error()),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// AuthHandler Verify the account and return the token.
// After getting the oauth2 token, you can do things like register the application.
func AuthHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	verify, err := NewVerify(username).Verify(password)
	if err != nil {
		c.JSON(401, gin.H{
			"status": fmt.Sprintf("VERIFY ACCOUNTS ERROR: %v", err.Error()),
		})
		return
	}
	// generator token.
	var (
		issuer = config.GetDomain()
		expir  = config.GetAuthTokenExpir()
		secret = config.GetAuthTokenSecret()
	)

	userdata := jwt.NewUserdata(strconv.Itoa(int(verify.ID)), verify.Username, verify.Mail)
	claims := jwt.NewRegisteredClaims(issuer, "", "", expir)
	generator, err := jwt.NewClaims(userdata, claims).JWTTokenGenerator(secret)

	c.JSON(200, gin.H{
		"status":       "ok",
		"access_token": generator,
	})
}

func NewApplicationHandler(c *gin.Context) {
	name := c.PostForm("name")
	homepageUrl := c.PostForm("homepage_url")
	description := c.PostForm("description")
	authorizationCallbackURL := c.PostForm("authorization_callback_url")

	n := NewOauthClients(GetUserId(c), name, homepageUrl, description, authorizationCallbackURL)
	create, err := n.Create()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"name":                       create.Name,
		"homepage_url":               create.HomepageURL,
		"description":                create.Description,
		"authorization_callback_url": create.AuthorizationCallbackURL,
		"client_id":                  strconv.Itoa(int(create.Model.ID)),
		"client_secrets":             create.Secrets,
	})
}
