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

func TestListAssets(t *testing.T) {
	data, err := os.Open("testdata/assets.json")
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
	assets, err := c.ListAssets(context.Background(), withURL(hs.URL))
	assert.NoError(t, err)
	assert.NotNil(t, assets)
	assert.Len(t, assets, 1031)
	assert.Equal(t, uint64(3971498809), c.Remaining())
	assert.Equal(t, uint64(0), c.RemainingPaid())
}

func TestGetAssetBySymbol(t *testing.T) {
	data, err := os.Open("testdata/asset.json")
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
	asset, err := c.GetAssetBySymbol(context.Background(), "btc", withURL(hs.URL))
	assert.NoError(t, err)
	assert.NotNil(t, asset)
	assert.Equal(t, uint64(3001498809), c.Remaining())
	assert.Equal(t, uint64(3232), c.RemainingPaid())
}
