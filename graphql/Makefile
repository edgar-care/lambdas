all: install

build:
	@go build main.go

clean:
	@rm -rf main

install:
	@go mod tidy

start:
	@go run main.go


.PHONY: all \
		build \
		install \
		start
