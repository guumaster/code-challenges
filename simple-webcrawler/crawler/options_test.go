package crawler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithHTTPClient(t *testing.T) {
	c := &Crawler{}
	client := &http.Client{}

	WithHTTPClient(client)(c)

	require.Equal(t, client, c.client)
}

func TestWithMaxThreads(t *testing.T) {
	c := &Crawler{}
	maxThreads := 10

	WithMaxThreads(maxThreads)(c)

	require.NotNil(t, c.queue)
}
