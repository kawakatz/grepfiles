# Reference: https://github.com/projectdiscovery/httpx/blob/master/Makefile
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY=bin/grepfiles
GOBIN=~/go/bin
    
all: build
build:
		$(GOBUILD) -v -o "bin/grepfiles" cmd/grepfiles/grepfiles.go
test: 
		$(GOTEST) -v ./...
tidy:
		$(GOMOD) tidy
install:	$(BINARY)
		install -s $(BINARY) $(GOBIN)