package cryptowatch

import (
	"context"
	"strings"
)

// PairDetail holds details of a pair.
type PairDetail struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Base   struct {
		ID     int    `json:"id"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
		Fiat   bool   `json:"fiat"`
		Route  string `json:"route"`
	} `json:"base"`
	Quote struct {
		ID     int    `json:"id"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
		Fiat   bool   `json:"fiat"`
		Route  string `json:"route"`
	} `json:"quote"`
	Route   string   `json:"route"`
	Markets []Market `json:"markets"`
}

// GetPairDetail returns the detail of a pair.
func (c *Client) GetPairDetail(ctx context.Context, pair string, opts ...Option) (PairDetail, error) {
	var as PairDetail
	opts = append([]Option{withURL(strings.Join([]string{host, "pairs", pair}, "/"))}, opts...)
	return as, c.doRequest(ctx, &as, opts...)
}
