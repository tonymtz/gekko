package oauth2

import (
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestProvider_RedirectUrl(t *testing.T) {
	// Setup

	leProvider := NewProvider(
		"my_key",
		"my_secret",
		"my_redirect",
		"http://sample.com/oauth2/authorize?client_id=%v&response_type=code&redirect_uri=%v",
		"nothing_here",
	)

	// Case

	url := leProvider.RedirectUrl()
	expectedUrl := "http://sample.com/oauth2/authorize?client_id=my_key&response_type=code&redirect_uri=my_redirect"

	assert.Equal(t, url, expectedUrl,
		fmt.Sprintf("Expected same value, got %v, want %v", url, expectedUrl),
	)
}

func TestProvider_ExchangeToken(t *testing.T) {
	// Setup
	var sampleResponse = `{
		"access_token": "my_unique_token",
		"token_type": "bearer",
		"uid": "12345"
	}`

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, sampleResponse)
	}))

	defer testServer.Close()

	leProvider := NewProvider(
		"my_key",
		"my_secret",
		"my_redirect",
		"nothing_here",
		testServer.URL,
	)

	// Case

	token, err := leProvider.ExchangeToken("my_code")

	if assert.Nil(t, err) {
		c := token.Token
		d := "my_unique_token"

		assert.Equal(t, c, d,
			fmt.Sprintf("Expected same value, got %v, want %v", c, d),
		)
	}
}
