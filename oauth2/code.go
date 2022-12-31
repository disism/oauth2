package oauth2

type Code interface {
	Create() error
}

func NewOauthAuthorizationCodes(userId uint, clientId string, code string) *OauthAuthorizationCodes {
	return &OauthAuthorizationCodes{UserId: userId, ClientId: clientId, Code: code}
}
