EXECUTABLE=crawler
VERSION=$(shell git describe --tags --always --long)
WINDOWS=$(EXECUTABLE)_windows.exe
LINUX=$(EXECUTABLE)_linux
DARWIN=$(EXECUTABLE)_darwin

.PHONY: all test clean

all: test clean build

test:
	go test -cover ./...

build: windows linux darwin
	@echo version: $(VERSION)

windows: $(WINDOWS)

linux: $(LINUX)

darwin: $(DARWIN)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o bin/$(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/main.go
$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o bin/$(LINUX) -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o bin/$(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/main.go

clean:
	rm -f bin/*