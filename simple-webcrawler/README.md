# Simple WebCrawler

This is a simple web crawler written in Golang. It visits each URL it finds on the same domain and prints each visited URL, as well as a list of links found on that page. The crawler is limited to one subdomain, so it does not follow external links. The output shows all links and marks `[new]` and `[external]` accordinly.


## Usage
You can run the crawler with a precompiled binary or directly from source. 


### From precompiled binary

The `bin` folder contains precompiled binaries for Windows, Mac and Linux. Choose the one matching your platform an run it with this command:

```sh
# On Linux
$/path/to/crawler> ./bin/crawler_linux <STARTING_URL> 

# On Window
C:\path\to\crawler>  .\bin\crawler_windows.exe <STARTING_URL> 
```

*NOTE:* Only Linux and Windows binary were tested during development. 


### From source

If you already have Go installed, you can run the crawler with the following command:

```bash
$/path/to/crawler> go run ./cmd/main.go <STARTING_URL>
```

## Dependencies

These are the libraries used for this crawler:


- `github.com/stretchr/testify`  a test toolkit with assertions and mocks.
- `github.com/guumaster/cligger` a simple CLI logger with colored symbols Includes fallbacks for Windows CMD.
- `golang.org/x/net`             a supplementary Go networking libraries for HTML tokenizer and parser.



## Testing

The crawler has comprehensive unit tests that use the `testify` library. To run the tests, use the following command:

```sh
go test -cover ./...

# Sample output:
# ?       github.com/guumaster/crawler/cmd        [no test files]
# ok      github.com/guumaster/crawler/crawler    0.287s  coverage: 92.1% of statements
# ok      github.com/guumaster/crawler/extractor  0.005s  coverage: 96.0% of statements
# ok      github.com/guumaster/crawler/queue      0.005s  coverage: 100.0% of statements

```

## Build

You can build the project using the `Makefile` available with the following command: 

```sh
$> make all
```

It will run tests and create new binaries for Linux, Windows and MacOS.



## Developer Notes

See [DEV Notes file](./DEV_NOTES.md)
