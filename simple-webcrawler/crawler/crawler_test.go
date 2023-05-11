package crawler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/guumaster/crawler/crawler"
)

type SpyErrorCallback struct {
	mock.Mock
}

func (e *SpyErrorCallback) Call(_ error) {
	e.Called()
}

type SpyRequestCallback struct {
	mock.Mock
}

func (e *SpyRequestCallback) Call(_ *http.Request) {
	e.Called()
}

type testRoundTripper struct{ mock.Mock }

func (rt *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	args := rt.Called(req)
	rsp, _ := args.Get(0).(*http.Response)
	return rsp, args.Error(1)
}

func TestNewCrawler(t *testing.T) {
	c := crawler.NewCrawler()
	require.NotNil(t, c)
}

func TestVisit(t *testing.T) {
	// Start a mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	response := &http.Response{StatusCode: http.StatusOK}

	// Create a mock HTTP client that returns the response from the mock server
	var transport testRoundTripper
	transport.On("RoundTrip", mock.Anything).
		Return(response, nil)
	defer transport.AssertExpectations(t)

	// Create a crawler instance with the mock HTTP client
	c := crawler.NewCrawler(
		crawler.WithHTTPClient(&http.Client{Transport: &transport}),
		crawler.WithMaxThreads(1),
	)

	// Visit the mock server
	err := c.Visit(ts.URL)
	require.NoError(t, err)

	c.Wait()
}

func TestVisitOnHTML(t *testing.T) {
	expectedHTML := "<html><head><title>Test Page</title></head><body><h1>Hello, World!</h1></body></html>"

	// Start a mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedHTML)
	}))
	defer ts.Close()

	// Create a new crawler instance
	c := crawler.NewCrawler()

	// Define a mock callback function to test OnHTML
	var responseBody string
	// Set the mock callback function as the HTML callback for the crawler
	c.OnHTML(func(url, html string) {
		responseBody = html
	})

	// Visit the mock server
	err := c.Visit(ts.URL)
	require.NoError(t, err)

	// Wait for all fetches to finish
	c.Wait()

	// Assert that the response body contains the expected HTML
	require.Equal(t, expectedHTML, responseBody)
}

func TestVisitWithError(t *testing.T) {
	// Start a mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	var transport testRoundTripper
	httpErr := errors.New("HTTP error")
	transport.
		On("RoundTrip", mock.Anything).
		Return(nil, httpErr)

	defer transport.AssertExpectations(t)

	c := crawler.NewCrawler(crawler.WithHTTPClient(&http.Client{Transport: &transport}))

	c.OnError(func(err error) {
		require.Error(t, err)
		require.Contains(t, err.Error(), "HTTP error")
	})
	// Visit the mock server
	err := c.Visit(ts.URL)
	require.NoError(t, err)

	c.Wait()
}
func TestVisitWithURLParsingError(t *testing.T) {
	c := crawler.NewCrawler()

	// Visit an invalid URL to trigger an error
	err := c.Visit("http://invalid_url.com/%")
	require.Error(t, err)

	c.Wait()
}
func TestVisitWithEmptyURL(t *testing.T) {
	c := crawler.NewCrawler()

	// check for empty URL to trigger an error
	err := c.Visit("")
	require.Error(t, err)

	c.Wait()
}

func TestVisitWithUnreachableURL(t *testing.T) {
	c := crawler.NewCrawler()

	spyCallback := new(SpyErrorCallback)

	spyCallback.On("Call").Return()

	c.OnError(spyCallback.Call)

	// check for empty URL to trigger an error
	err := c.Visit("http://un.reach.ab.le")
	require.NoError(t, err)

	c.Wait()

	spyCallback.AssertCalled(t, "Call", mock.Anything)
}

func TestVisitOnRequestCallback(t *testing.T) {
	c := crawler.NewCrawler()

	spy := new(SpyRequestCallback)
	spy.On("Call").Return()

	c.OnRequest(spy.Call)

	// check for empty URL to trigger an error
	err := c.Visit("http://un.reach.ab.le")
	require.NoError(t, err)

	c.Wait()

	spy.AssertCalled(t, "Call", mock.Anything)
}

func TestVisitHasVisited(t *testing.T) {
	// Start a mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	response := &http.Response{StatusCode: http.StatusOK}

	// Create a mock HTTP client that returns the response from the mock server once
	var transport testRoundTripper
	transport.On("RoundTrip", mock.Anything).
		Return(response, nil).
		Once()
	defer transport.AssertExpectations(t)

	// Create a crawler instance with the mock HTTP client
	mockClient := crawler.WithHTTPClient(&http.Client{Transport: &transport})
	c := crawler.NewCrawler(mockClient)

	// check for empty URL to trigger an error
	err := c.Visit("http://example.com")
	require.NoError(t, err)

	err = c.Visit("http://example.com")
	require.NoError(t, err)

	c.Wait()
}
