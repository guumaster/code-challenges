
## Developer Notes

As requested in the excersize, the crawler does not use any existing crawler library. I wanted to have a robust crawler that could be reused for other tasks without changing much, so that's why I've chosen to use a similar interface to `go-colly` library with a basic implementation. This not only separates the link extractor logic from the crawler but also allows to do a quick drop-in replacement of this crawler with the `go-colly` library with little code changes.

In order to code this exercise with the time available, I made some decissions that I consider worth mentioning:

- the CLI program accepts only the starting URL and have good defaults parameters. I've decided not to add any extra flags leaving it for possible improvements to open parameters like timeout and max number of threads with the standard library or with `cobra` package for more complex projects. 

- it only runs GET requests and extract from HTML body. 

- it may not follow redirects properly.

- it doesn't handle retries on failed request, only shows an error message.

- it's designed to run several processes concurrently making use of goroutines and channels to control fetching child pages in parallel with a limit on the max number of concurrent processes.

- it doesn't have any restriction of the depth level of requests due to the implied extra code needed to change the current implementation.

- in this implementation event handlers `OnError`, `OnHTML` and `OnRequest` only accepts a single handler. Code changes are needed to work with a list of handlers instead.

- crawler does not use a cancellable context, instead of `context.Background()` due to necessary code changes to sync channels and contexts.

