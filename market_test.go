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

func TestListMarkets(t *testing.T) {
	data, err := os.Open("testdata/markets.json")
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
	markets, err := c.ListMarkets(context.Background(), withURL(hs.URL))
	assert.NoError(t, err)
	assert.NotNil(t, markets)
	assert.Equal(t, uint64(3971498809), c.Remaining())
	assert.Equal(t, uint64(0), c.RemainingPaid())
}

func TestListMarketDetails(t *testing.T) {
	data, err := os.Open("testdata/market_details.json")
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
	market, err := c.ListMarketDetails(context.Background(), "kraken", "btceur", withURL(hs.URL))
	assert.NoError(t, err)
	assert.NotNil(t, market)
	assert.Equal(t, uint64(3971498809), c.Remaining())
	assert.Equal(t, uint64(0), c.RemainingPaid())
}
