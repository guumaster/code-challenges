package extractor

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type PageLinks struct {
	URL           string
	BaseURL       string
	InternalLinks []string
	ExternalLinks []string
}

// Collect all links from response body and return it as an array of strings.
func GetAllLinks(srcURL, body string) (*PageLinks, error) {
	baseURL, err := getBaseURL(srcURL)
	if err != nil {
		return nil, err
	}

	links := readAllLinks(baseURL, body)
	internal, external := splitLinks(baseURL, links)

	return &PageLinks{
		URL:           srcURL,
		BaseURL:       baseURL,
		InternalLinks: unique(internal),
		ExternalLinks: unique(external),
	}, nil
}

func getBaseURL(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	u.Path = ""

	return u.String(), nil
}

func readAllLinks(baseURL, body string) []string {
	r := strings.NewReader(body)
	var links []string
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link := formatURL(baseURL, attr.Val)
						if link != "" {
							links = append(links, link)
						}
					}
				}
			}
		}
	}
}

func formatURL(base string, l string) string {
	base = strings.TrimSuffix(base, "/")

	formatted := ""
	switch {
	case strings.HasPrefix(l, "https://"):
		fallthrough
	case strings.HasPrefix(l, "http://"):
		formatted = l
	case strings.HasPrefix(l, "/"):
		formatted = base + l
	}

	u, err := url.Parse(formatted)
	if err != nil {
		return ""
	}
	u.Fragment = ""

	return u.String()
}

func splitLinks(baseURL string, links []string) ([]string, []string) {
	internal := make([]string, 0)
	external := make([]string, 0)

	for _, link := range links {
		if strings.HasPrefix(link, baseURL) {
			internal = append(internal, link)
		} else {
			external = append(external, link)
		}
	}

	return internal, external
}

func unique(strs []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strs {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
