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

func TestGetPrice(t *testing.T) {
	data, err := os.Open("testdata/price.json")
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
	price, err := c.GetPrice(context.Background(), "kraken", "btceur", withURL(hs.URL))
	assert.NoError(t, err)
	assert.Equal(t, float64(6595), price)
	assert.Equal(t, uint64(3971498809), c.Remaining())
	assert.Equal(t, uint64(0), c.RemainingPaid())
}
