package cryptowatch

import (
	"net/http"
	"sync/atomic"
	"time"
)

const (
	defaultTimeout = 10
	host           = "https://api.cryptowat.ch"
)

// Client for cryptowatch API.
type Client struct {
	c             *http.Client
	remaining     *uint64
	remainingPaid *uint64
}

// NewClient returns a new instance of Client.
func NewClient() *Client {
	return &Client{
		c: &http.Client{
			Timeout: defaultTimeout * time.Second,
		},
		remaining:     new(uint64),
		remainingPaid: new(uint64),
	}
}

// Remaining returns the remaining counter.
func (c *Client) Remaining() uint64 {
	return atomic.LoadUint64(c.remaining)
}

// RemainingPaid returns the remaining paid counter.
func (c *Client) RemainingPaid() uint64 {
	return atomic.LoadUint64(c.remainingPaid)
}
