package cryptowatch

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListExchanges(t *testing.T) {
	data, err := os.Open("testdata/exchanges.json")
	if err != nil {
		t.Error(err)
	}
	defer data.Close()

	hs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(w, data)
	}))
	defer hs.Close()

	c := NewClient()
	exchanges, err := c.ListExchanges(context.Background(), withURL(hs.URL))
	assert.NoError(t, err)
	assert.NotNil(t, exchanges)
	assert.Equal(t, uint64(3971498809), c.Remaining())
	assert.Equal(t, uint64(0), c.RemainingPaid())
}

func TestListExchangeDetails(t *testing.T) {
	data, err := os.Open("testdata/exchange.json")
	if err != nil {
		t.Error(err)
	}
	defer data.Close()

	hs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(w, data)
	}))
	defer hs.Close()

	c := NewClient()
	exchange, err := c.ListExchangeDetails(context.Background(), "kraken", withURL(hs.URL))
	assert.NoError(t, err)
	assert.NotNil(t, exchange)
	assert.Equal(t, uint64(3971498809), c.Remaining())
	assert.Equal(t, uint64(0), c.RemainingPaid())
}

func TestListExchangeMarkets(t *testing.T) {
	data, err := os.Open("testdata/market.json")
	if err != nil {
		t.Error(err)
	}
	defer data.Close()

	hs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(w, data)
	}))
	defer hs.Close()

	c := NewClient()
	market, err := c.ListExchangeMarkets(context.Background(), "kraken", withURL(hs.URL))
	assert.NoError(t, err)
	assert.NotNil(t, market)
	assert.Equal(t, uint64(3971498809), c.Remaining())
	assert.Equal(t, uint64(0), c.RemainingPaid())
}
