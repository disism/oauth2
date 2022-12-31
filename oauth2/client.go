package oauth2

import (
	"crypto/rand"
	"fmt"
	"gorm.io/gorm"
)

type Client interface {
	Create() (*OauthClients, error)
	Get() (*OauthClients, error)
}

func NewOauthClients(userId uint, name string, homepageURL string, description string, authorizationCallbackURL string) *OauthClients {
	return &OauthClients{UserId: userId, Name: name, HomepageURL: homepageURL, Description: description, AuthorizationCallbackURL: authorizationCallbackURL}
}

func (a *OauthClients) Create() (*OauthClients, error) {
	if err := db.AutoMigrate(&OauthClients{}); err != nil {
		return nil, err
	}

	a.Secrets = RandSecrets(24)
	if err := db.Debug().Create(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func NewClientId(clientId uint) *OauthClients {
	return &OauthClients{
		Model: gorm.Model{
			ID: clientId,
		},
	}
}
func (a *OauthClients) Get() (*OauthClients, error) {
	if err := db.Debug().
		Table("oauth_clients").
		Where("id = ?", a.ID).
		First(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func RandSecrets(num int) string {
	b := make([]byte, num)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
