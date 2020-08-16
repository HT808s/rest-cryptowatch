package cryptowatch

import (
	"context"
	"strings"
)

type price struct {
	Price float64 `json:"price"`
}

// GetPrice returns the price of a pair in the given exchange.
func (c *Client) GetPrice(ctx context.Context, exchange, pair string, opts ...Option) (float64, error) {
	var as price
	opts = append([]Option{withURL(strings.Join([]string{host, "markets", exchange, pair, "price"}, "/"))}, opts...)
	return as.Price, c.doRequest(ctx, &as, opts...)
}
