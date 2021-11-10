package front

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const TokenURL = "https://app.frontapp.com/oauth/token"

// WithCredentials oauth2 HTTP client.
func WithCredentials(id, secret string) ClientOption {
	config := &clientcredentials.Config{
		ClientID:     id,
		ClientSecret: secret,
		TokenURL:     TokenURL,
		AuthStyle:    oauth2.AuthStyleInHeader,
	}

	return WithHTTPClient(config.Client(context.TODO()))
}

// StringParam creates a string pointer for optional params.
func StringParam(v string) *string {
	return &v
}

// BooleanParam creates a boolean pointer for optional boolean params.
func BooleanParam(v bool) *bool {
	return &v
}
