package cryptowatch

import (
	"net/http"
	"net/url"
)

// Option is an optional argument used to add data to a request.
type Option interface {
	Apply(r *http.Request, q url.Values)
}

type optAPIKey struct {
	secret string
}

// WithAPIKey sets apikey to header of the request
func WithAPIKey(s string) Option {
	return &optAPIKey{
		secret: s,
	}
}

func (o *optAPIKey) Apply(r *http.Request, q url.Values) {
	const apikeyHeader = "X-CW-API-Key"
	r.Header.Set(apikeyHeader, o.secret)
}

type optURL struct {
	url string
}

// withURL sets url to the request
func withURL(s string) Option {
	return &optURL{
		url: s,
	}
}

func (o *optURL) Apply(r *http.Request, q url.Values) {
	r.URL, _ = r.URL.Parse(o.url)
}
