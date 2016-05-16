package provider

import "github.com/tonymtz/gekko/services/oauth2"

const (
	DROPBOX_AUTHORIZATION_URL = "https://accounts.google.com/o/oauth2/v2/auth?client_id=%v&response_type=code&redirect_uri=%v&scope=email profile"
	DROPBOX_TOKEN_EXCHANGE_URL = "https://www.googleapis.com/oauth2/v4/token"
)

type Dropbox struct {
	oauth2.Provider
}

func NewDropbox(key, secret, redirectUrl string) *Dropbox {
	return &Dropbox{
		oauth2.NewProvider(
			key,
			secret,
			redirectUrl,
			DROPBOX_AUTHORIZATION_URL,
			DROPBOX_TOKEN_EXCHANGE_URL,
		),
	}
}
