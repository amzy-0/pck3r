.PHONY: all build clean deps

MODULE_NAME := pck3r
BINARY_NAME := pck3r

all: build

deps:
	go get github.com/fatih/color

build: deps
	go build -o $(BINARY_NAME) main.go

clean:
	rm -f $(BINARY_NAME)

install: build
	sudo install -m 755 $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
