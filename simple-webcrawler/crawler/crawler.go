package crawler

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/guumaster/crawler/queue"
)

type Crawler struct {
	client          *http.Client
	context         context.Context
	wg              *sync.WaitGroup
	visitedUrls     sync.Map
	htmlCallback    OnHTMLCallback
	requestCallback OnRequestCallback
	errorCallback   OnErrorCallback
	queue           *queue.Queue
}

type OnRequestCallback func(req *http.Request)

type OnHTMLCallback func(url, html string)

type OnErrorCallback func(err error)

const defaultRequestTimeout = 10

var ErrEmptyURL = errors.New("url can't be empty")

func NewCrawler(opts ...Option) *Crawler {
	client := &http.Client{
		Timeout: time.Second * defaultRequestTimeout,
	}

	c := &Crawler{
		client:  client,
		context: context.Background(),
		wg:      &sync.WaitGroup{},
	}

	for _, opt := range opts {
		opt(c)
	}
	if c.queue != nil {
		c.queue.Start()
	}

	return c
}

func (c *Crawler) Visit(base string) error {
	u, err := parseURL(base)
	if err != nil {
		return err
	}

	if c.HasVisited(u) {
		return nil
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	fetchFn := func() {
		req = req.WithContext(c.context)

		c.visitedUrls.Store(u, true)

		c.wg.Add(1)
		c.fetch(u, req)
	}

	if c.queue != nil {
		c.queue.AddJob(fetchFn)
		return nil
	}

	fetchFn()
	return nil
}

func (c *Crawler) Wait() {
	c.wg.Wait()
	if c.queue != nil {
		c.queue.Wait()
	}
}

func (c *Crawler) HasVisited(url string) bool {
	_, ok := c.visitedUrls.Load(url)
	return ok
}

func (c *Crawler) OnError(f OnErrorCallback) {
	c.errorCallback = f
}

func (c *Crawler) OnHTML(f OnHTMLCallback) {
	c.htmlCallback = f
}

func (c *Crawler) OnRequest(f OnRequestCallback) {
	c.requestCallback = f
}

func (c *Crawler) handleOnError(err error) {
	if c.errorCallback != nil {
		c.errorCallback(err)
	}
}

func (c *Crawler) handleOnHTML(url, body string) {
	if c.htmlCallback != nil {
		c.htmlCallback(url, body)
	}
}

func (c *Crawler) handleOnRequest(req *http.Request) {
	if c.requestCallback != nil {
		c.requestCallback(req)
	}
}

func (c *Crawler) fetch(url string, req *http.Request) {
	defer c.wg.Done()

	c.handleOnRequest(req)

	resp, err := c.client.Do(req)
	if err != nil {
		c.handleOnError(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.handleOnError(err)
		return
	}
	c.handleOnHTML(url, string(body))
}

func parseURL(u string) (string, error) {
	if u == "" {
		return "", ErrEmptyURL
	}

	parsedURL, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	return parsedURL.String(), nil
}
