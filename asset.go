package cryptowatch

import "context"

// Asset is something that is traded, like a crypto or fiat currency.
type Asset struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Fiat   bool   `json:"fiat"`
	Route  string `json:"route"`
}

// AssetDetail holds information about an asset.
type AssetDetail struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	// Fiat is true when the asset is fiat currency (such as usd, and false when it is a cryptocurrency.
	Fiat    bool `json:"fiat"`
	Markets struct {
		Base []struct {
			ID       int    `json:"id"`
			Exchange string `json:"exchange"`
			Pair     string `json:"pair"`
			Active   bool   `json:"active"`
			Route    string `json:"route"`
		} `json:"base"`
		Quote []struct {
			ID       int    `json:"id"`
			Exchange string `json:"exchange"`
			Pair     string `json:"pair"`
			Active   bool   `json:"active"`
			Route    string `json:"route"`
		} `json:"quote"`
	} `json:"markets"`
}

// ListAssets returns all assets.
func (c *Client) ListAssets(ctx context.Context, opts ...Option) ([]Asset, error) {
	var as []Asset
	return as, c.doRequest(ctx, &as, append([]Option{withURL(host + `/assets`)}, opts...)...)
}

// GetAssetBySymbol returns an asset details.
func (c *Client) GetAssetBySymbol(ctx context.Context, symbol string, opts ...Option) (AssetDetail, error) {
	var as AssetDetail
	return as, c.doRequest(ctx, &as, append([]Option{withURL(host + `/assets/` + symbol)}, opts...)...)
}
