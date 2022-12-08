BINARY_NAME     = ping
GO_BINARY_PATH  =$(shell go env GOPATH)/bin

all: install

build:
		@go build -o $(BINARY_NAME)

install:
		@go mod tidy


.PHONY: all \
			build \
			install
