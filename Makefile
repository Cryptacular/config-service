all: deps build test

ci: deps ci-build test

build:
	go build

ci-build:
	go build -o index.config-service

test:
	go test -v

format:
	go fmt

deps:
	go get -t ./...