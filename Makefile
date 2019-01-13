# vi: ft=make

GOPATH:=$(shell go env GOPATH)

build:
	go build -o protoc-gen-swiftgrpcrx main.go

install:
	sudo cp protoc-gen-swiftgrpcrx /usr/local/bin

test:
	go test -p 1 -race -cover -v ./...
