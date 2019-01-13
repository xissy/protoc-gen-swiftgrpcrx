# vi: ft=make

GOPATH:=$(shell go env GOPATH)

build:
	go build -o protoc-gen-swiftgrpcrx main.go

install:
	go install
	
test:
	go test -p 1 -race -cover -v ./...
