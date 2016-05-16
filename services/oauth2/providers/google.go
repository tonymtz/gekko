package providers

import "github.com/tonymtz/gekko/services/oauth2"

const (
	GOOGLE_AUTHORIZATION_URL = "https://accounts.google.com/o/oauth2/v2/auth?client_id=%v&response_type=code&redirect_uri=%v&scope=email profile"
	GOOGLE_TOKEN_EXCHANGE_URL = "https://www.googleapis.com/oauth2/v4/token"
)

func NewGoogle(key, secret, redirectUrl string) oauth2.IProvider {
	return oauth2.NewProvider(
		key,
		secret,
		redirectUrl,
		GOOGLE_AUTHORIZATION_URL,
		GOOGLE_TOKEN_EXCHANGE_URL,
	)
}
