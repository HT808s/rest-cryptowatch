package cryptowatch

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sync/atomic"
)

// Sends request to API
// Apply option to request
// Handle response whether it is an error or not
// Updates the remaining and remainingPaid counters
func (c *Client) doRequest(ctx context.Context, data interface{}, opts ...Option) error {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)
	if err != nil {
		return err
	}

	q := r.URL.Query()
	for _, o := range opts {
		if o == nil {
			continue
		}
		o.Apply(r, q)
	}

	r.URL.RawQuery = q.Encode()

	resp, err := c.c.Do(r)
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		var e errResponse
		if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return err
		}

		atomic.StoreUint64(c.remaining, e.Allowance.Remaining)
		atomic.StoreUint64(c.remainingPaid, e.Allowance.RemainingPaid)
		return errors.New(e.Error)
	}

	rawMap := make(map[string]json.RawMessage)
	if err = json.NewDecoder(resp.Body).Decode(&rawMap); err != nil {
		return err
	}

	if raw, ok := rawMap["allowance"]; ok {
		var all allowance
		if err := json.Unmarshal(raw, &all); err != nil {
			return err
		}
		atomic.StoreUint64(c.remaining, all.Remaining)
		atomic.StoreUint64(c.remainingPaid, all.RemainingPaid)
	}

	if raw, ok := rawMap["result"]; ok {
		return json.Unmarshal(raw, data)
	}

	return errors.New("internal error")
}
