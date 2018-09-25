# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=infrabot

all: test build install

build:
	mkdir -p bin
	$(GOBUILD) -o bin/$(BINARY_NAME) -v

install:
	cp bin/$(BINARY_NAME) /usr/local/bin/

test:
	echo "TODO"

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

deps:
	$(GOGET) github.com/urfave/cli
