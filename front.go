package front

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/time/rate"
)

type transport struct {
	bearer    string
	limiter   *rate.Limiter
	transport http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := t.limiter.Wait(req.Context()); err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+t.bearer)
	return t.transport.RoundTrip(req)
}

// WithFrontClient creates a Front compatible client.
//
// It handles injection of the Front bearer token, API backoff requests and rate limits.
// See https://dev.frontapp.com/docs/rate-limiting
func WithFrontClient(token string, limit float64) ClientOption {
	// The retryablehttp client is used to handle Retry-After headers.
	client := retryablehttp.NewClient()
	client.RetryMax = 1

	// Injects the authorization bearer token header into every request.
	client.HTTPClient.Transport = &transport{
		bearer:    token,
		limiter:   rate.NewLimiter(rate.Limit(limit), 1),
		transport: client.HTTPClient.Transport,
	}

	return WithHTTPClient(client.StandardClient())
}

// StringParam creates a string pointer for optional params.
func StringParam(v string) *string {
	return &v
}

// BooleanParam creates a boolean pointer for optional boolean params.
func BooleanParam(v bool) *bool {
	return &v
}
