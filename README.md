# Code challenges

This folder contains code examples that I've recently created for you to check and analize how I work. Here you have a quick summary, but then each project contains two files where you can get more info, one `README.md` with description about the code and usage, and a `DEV_NOTES.md` where you can read about my thoughts, decissions and tradeoff of each one.


### http-log-parser

The requirements for this exercise was to create a parser for csv files containing http log data and show some stats and insights about the data. I choose a very flexible design that allowed not only to collect stats and show logs but also is open for new functionality following simple interfaces.

This project is close to a year and half old, maybe some libraries may need review, but the binary and tests should work as expected. Also, probably the linter config is not up-to-date and needs review.


### Web Crawler

The requirement for this exercise was to create a simple web crawler cappable of extracting links from any page and follow links in the same subdomain but not external links. With an additional contrain of not using existing crawling libraries. So I choose to implement a minimun but similar interface as the known `go-colly` package, allowing the code to be flexible an well separated between extraction logic from crawling logic allowing to create different programs with the same basic crawler or just replace it the full `go-colly` library with a few code changes. 


