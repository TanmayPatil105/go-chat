BINARY_NAME = go-chat
SRCS = server.go
REPO = github.com/TanmayPatil105/go-chat

all: build

.PHONY: deps
deps:
	go get -d -v ${REPO}/...

.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: run
run:
	@go run ${SRCS}

.PHONY: build
build: deps
	go build -o ${BINARY_NAME} ${SRCS}

.PHONY: clean
clean:
	rm -rf ${BINARY_NAME}
