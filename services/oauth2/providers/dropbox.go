package provider

import "github.com/tonymtz/gekko/services/oauth2"

const (
	DROPBOX_AUTHORIZATION_URL = "https://www.dropbox.com/oauth2/authorize?client_id=%v&response_type=code&redirect_uri=%v"
	DROPBOX_TOKEN_EXCHANGE_URL = "https://api.dropboxapi.com/oauth2/token"
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
