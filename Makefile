all: build

.PHONY: build
build:
	go build -o go-chat server.go
