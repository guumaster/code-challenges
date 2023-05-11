package crawler

import (
	"net/http"
	"time"

	"github.com/guumaster/crawler/queue"
)

// Option wrapper function to contains configurable Crawler options.
type Option func(*Crawler)

func WithHTTPClient(client *http.Client) Option {
	return func(c *Crawler) {
		c.client = client
	}
}

func WithMaxThreads(maxThreads int) Option {
	return func(c *Crawler) {
		c.queue = queue.NewQueue(maxThreads)
	}
}
func WithRequestTimeout(t time.Duration) Option {
	return func(c *Crawler) {
		c.client.Timeout = t
	}
}
