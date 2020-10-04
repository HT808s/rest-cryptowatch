package cryptowatch

import (
	"context"
	"strings"
)

// Market is a pair listed on an exchange.
type Market struct {
	ID       int    `json:"id"`
	Exchange string `json:"exchange"`
	Pair     string `json:"pair"`
	Active   bool   `json:"active"`
	Route    string `json:"route"`
}

// ListMarkets returns markets associated with an exchange.
func (c *Client) ListMarkets(ctx context.Context, opts ...Option) ([]Market, error) {
	var as []Market
	opts = append([]Option{withURL(host + `/markets`)}, opts...)
	return as, c.doRequest(ctx, &as, opts...)
}

// MarketDetails holds information about the market.
type MarketDetails struct {
}

// ListMarketDetails returns markets associated with an exchange.
func (c *Client) ListMarketDetails(ctx context.Context, exchange, pair string, opts ...Option) (MarketDetails, error) {
	var as MarketDetails
	opts = append([]Option{withURL(strings.Join([]string{host, "markets", exchange, pair}, "/"))}, opts...)
	return as, c.doRequest(ctx, &as, opts...)
}
