package oauth2

import "gorm.io/gorm"

type OauthUsers struct {
	gorm.Model
	Username string `gorm:"index;type:text;username;unique" validate:"required,min=4,max=16"`
	Mail     string `gorm:"index;type:text;mail;unique" validate:"required,email"`
	Password string `gorm:"type:text;password" validate:"required,min=8,max=24"`
	Bio      string `gorm:"type:text;bio"`
	Avatar   string `gorm:"type:text;avatar"`
}

type OauthAccessTokens struct {
	gorm.Model

	UserId   uint   `gorm:"index;type:bigint;user_id"`
	ClientId string `gorm:"type:bigint;client_id"`
	Token    string `gorm:"type:text;token"`
}

type OauthRefreshTokens struct {
}

type OauthAuthorizationCodes struct {
	UserId   uint   `gorm:"index;type:bigint;user_id"`
	ClientId string `gorm:"type:bigint;client_id"`
	Code     string `gorm:"type:text;code"`
}

type OauthClients struct {
	gorm.Model
	UserId                   uint   `gorm:"index;type:bigint;user_id"`
	Name                     string `gorm:"type:text;name"`
	HomepageURL              string `gorm:"type:text;homepage_url"`
	Description              string `gorm:"type:text;description"`
	AuthorizationCallbackURL string `gorm:"type:text;authorization_callback_url"`
	Secrets                  string `gorm:"type:text;secrets"`
}

type OauthScopes struct {
}

type OauthRoles struct {
}
