package cryptowatch

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_doRequest_failure(t *testing.T) {
	apikey := "api-key"
	hs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, apikey, r.Header.Get("X-CW-API-Key"))
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(w, `{
			"error":"Exchange not found",
			"allowance":{
			  "cost":20198,
			  "remaining":3963870513,
			  "remainingPaid":0,
			  "upgrade":"Upgrade for a higher allowance, starting at $15/month for 16 seconds/hour. https://cryptowat.ch/pricing"
			}
		  }`)
	}))
	defer hs.Close()

	var d errResponse
	err := NewClient().doRequest(context.Background(), &d, withURL(hs.URL), WithAPIKey(apikey))
	assert.Error(t, err)
	assert.EqualError(t, err, "Exchange not found")
}
