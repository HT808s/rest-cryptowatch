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

func TestGetPairDetail(t *testing.T) {
	data, err := os.Open("testdata/pair_detail.json")
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
	market, err := c.GetPairDetail(context.Background(), "btceur", withURL(hs.URL))
	assert.NoError(t, err)
	assert.NotNil(t, market)
	assert.Equal(t, uint64(3971498809), c.Remaining())
	assert.Equal(t, uint64(0), c.RemainingPaid())
}
