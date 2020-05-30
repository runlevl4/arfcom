include .env

GOBASE=$(shell pwd)
MAKEFLAGS += --silent
TAG="k8sninja/arfcom:$(BUILD_TAG)"

build-mac:
	echo "***** building mac binary *****"
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/arfcom api/cmd/main.go
	echo "***** build complete *****"

build-linux:
	echo "***** building linux binary *****"
	GOOS=linux GOARCH=amd64 go build -o bin/linux/arfcom api/cmd/main.go
	echo "***** build complete *****"

build-all: build-mac build-linux

docker-build:	
	docker build -t $(TAG) .

docker-push:
	docker push $(TAG)


