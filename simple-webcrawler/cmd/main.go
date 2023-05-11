package main

import (
	"fmt"
	"os"
	"time"

	"github.com/guumaster/crawler/crawler"
	"github.com/guumaster/crawler/extractor"

	"github.com/guumaster/cligger"
)

const defaultMaxThreads = 20

func main() {
	// check arguments
	if len(os.Args) == 1 {
		cligger.Error("missing starting url to run crawler.")
		return
	}
	startingURL := os.Args[1]

	// prepare crawler with options
	c := crawler.NewCrawler(
		crawler.WithRequestTimeout(1*time.Minute),
		crawler.WithMaxThreads(defaultMaxThreads),
	)

	// setup callbacks
	// c.OnRequest(func(req *http.Request) {
	//   cligger.Infof("Visiting: %s", req.URL)
	// })
	c.OnError(func(err error) {
		cligger.Warning("ERROR: %s", err)
	})
	// handle each page and visit children pages
	c.OnHTML(collectAllLinks(c))

	cligger.Info("Starting crawler for %s.", startingURL)

	err := c.Visit(startingURL)
	if err != nil {
		cligger.Fatal("error visiting starting url %s", startingURL)
		return
	}
	c.Wait()

	cligger.Success("Crawler done.")
}

// collectAllLinks handle each visited page html, show found links, and send.
func collectAllLinks(c *crawler.Crawler) crawler.OnHTMLCallback {
	return func(current, body string) {
		pageLinks, err := extractor.GetAllLinks(current, body)
		if err != nil {
			cligger.Error("error extracting links")
			return
		}

		totalLinks := len(pageLinks.InternalLinks) + len(pageLinks.ExternalLinks)
		if totalLinks == 0 {
			cligger.Warning("No Links found on [%s]", current)
			return
		}

		foundLinks := prepareLinkOutput(c, *pageLinks)
		for _, link := range pageLinks.InternalLinks {
			if !c.HasVisited(link) {
				errVisit := c.Visit(link)
				if errVisit != nil {
					cligger.Error("error visiting child url %s: %s", link, errVisit.Error())
					return
				}
			}
		}

		cligger.Info("Links found on [%s]: \n%s ", current, foundLinks)
	}
}

func prepareLinkOutput(c *crawler.Crawler, pageLinks extractor.PageLinks) string {
	foundLinks := ""
	newLinks := ""
	for _, link := range pageLinks.InternalLinks {
		if !c.HasVisited(link) {
			newLinks += fmt.Sprintf("\t[new] %s\n", link)
		} else {
			foundLinks += fmt.Sprintf("\t      %s\n", link)
		}
	}
	for _, link := range pageLinks.ExternalLinks {
		foundLinks += fmt.Sprintf("\t[ext] %s\n", link)
	}

	return newLinks + foundLinks
}
