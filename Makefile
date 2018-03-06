all: build test

build:
	go build -o main

aws-build:
	GOOS=linux go build -o main
	zip deployment.zip main

deploy: aws-build
	aws lambda update-function-code --function-name ConfigService --zip-file fileb://deployment.zip

test:
	go test -v

format:
	go fmt

deps:
	go get -t ./...