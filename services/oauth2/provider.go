package oauth2

import (
	"fmt"
	"net/url"
	"encoding/json"
	"net/http"
)

type Token struct {
	UID   string        `json:"uid"`
	Token string        `json:"access_token"`
	Error *string       `json:"error"`
}

type IProvider interface {
	RedirectUrl() string
	ExchangeToken(string) (*Token, error)
}

func NewProvider(key, secret, redirectUrl, authURL, exchangeURL string) Provider {
	return Provider{
		Key: key,
		Secret: secret,
		RedirectURL: redirectUrl,
		authURL: authURL,
		exchangeURL: exchangeURL,
	}
}

type Provider struct {
	IProvider,
	Key         string
	Secret      string
	RedirectURL string
	exchangeURL string
	authURL     string
}

func (this *Provider) RedirectUrl() string {
	return fmt.Sprintf(this.authURL, this.Key, this.RedirectURL)
}

func (this *Provider) ExchangeToken(code string) (*Token, error) {
	data := url.Values{}

	data.Add("code", code)
	data.Add("grant_type", "authorization_code")
	data.Add("client_id", this.Key)
	data.Add("client_secret", this.Secret)
	data.Add("redirect_uri", this.RedirectURL)

	resp, err := http.PostForm(this.exchangeURL, data)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	token := &Token{}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&token)

	return token, nil
}
