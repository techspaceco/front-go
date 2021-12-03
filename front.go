package front

import (
	"context"
	"net"
	"net/http"
	"time"

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

type transport struct {
	bearer    string
	transport http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// start := time.Now()
	req.Header.Set("Authorization", "Bearer "+t.bearer)
	res, err := t.transport.RoundTrip(req)
	if err != nil {
		// TODO: Logging?
		// log.Printf("ERR %s %s\n", req.URL.String(), err)
		return res, err
	}

	// TODO: Logging?
	// log.Printf("%d %s %s\n", res.StatusCode, req.URL.String(), time.Since(start))
	return res, err
}

// With bearer token.
func WithAuthorizationToken(token string) ClientOption {
	client := &http.Client{
		Transport: &transport{
			bearer: token,
			transport: &http.Transport{
				Dial: (&net.Dialer{Timeout: 5 * time.Second}).Dial,
			},
		},
	}

	return WithHTTPClient(client)
}

// StringParam creates a string pointer for optional params.
func StringParam(v string) *string {
	return &v
}

// BooleanParam creates a boolean pointer for optional boolean params.
func BooleanParam(v bool) *bool {
	return &v
}
