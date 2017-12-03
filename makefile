install:
	go get -d -v ./...
	go install -v ./...

build:
	go build

all: install build