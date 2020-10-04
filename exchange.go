package cryptowatch

import "context"

// List of supported exchanges.
//nolint:varcheck,deadcode,unused
const (
	Bitfinex      = "bitfinex"
	CoinbasePro   = "coinbase-pro"
	Bitstamp      = "bitstamp"
	Kraken        = "kraken"
	CEXIO         = "cexio"
	Gemini        = "gemini"
	Quoine        = "quoine"
	Liquid        = "liquid"
	bitFlyer      = "bitflyer"
	OKCoin        = "okcoin"
	BitMEX        = "bitmex"
	Huobi         = "huobi"
	Luno          = "luno"
	Poloniex      = "poloniex"
	Bisq          = "bisq"
	Bithumb       = "bithumb"
	Bittrex       = "bittrex"
	Binance       = "binance"
	BitBay        = "bitbay"
	Okex          = "okex"
	Coinone       = "coinone"
	HitBTC        = "hitbtc"
	BitZ          = "bitz"
	Gateio        = "gateio"
	BinanceUS     = "binance-us"
	KrakenFutures = "kraken-futures"
	DEXAggregated = "dex-aggregated"
	FTX           = "ftx"
	Deribit       = "deribit"
)

// Exchange is where all the action happens!
type Exchange struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Route  string `json:"route"`
	Active bool   `json:"active"`
}

// ExchangeDetail holds detailed information about an exhange.
type ExchangeDetail struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Routes struct {
		Markets string `json:"markets"`
	} `json:"routes"`
}

// ListExchanges list all exchanges.
func (c *Client) ListExchanges(ctx context.Context, opts ...Option) ([]Exchange, error) {
	var as []Exchange
	opts = append([]Option{withURL(host + "/exchanges")}, opts...)
	return as, c.doRequest(ctx, &as, opts...)
}

// ListExchangeDetails returns details of given exchange.
func (c *Client) ListExchangeDetails(ctx context.Context, exchange string, opts ...Option) (ExchangeDetail, error) {
	var as ExchangeDetail
	opts = append([]Option{withURL(host + "/exchanges/" + exchange)}, opts...)
	return as, c.doRequest(ctx, &as, opts...)
}

// ListExchangeMarkets returns markets associated with an exchange.
func (c *Client) ListExchangeMarkets(ctx context.Context, exchange string, opts ...Option) ([]Market, error) {
	var as []Market
	opts = append([]Option{withURL(host + "/markets/" + exchange)}, opts...)
	return as, c.doRequest(ctx, &as, opts...)
}
