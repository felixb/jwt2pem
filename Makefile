.PHONY: all build clean test

BIN_DIR := build
BIN_NAME := jwt2pem
BINS := $(BIN_DIR)/$(BIN_NAME)-linux-amd64 $(BIN_DIR)/$(BIN_NAME)-darwin-amd64 $(BIN_DIR)/$(BIN_NAME)-windows-amd64.exe $(BIN_DIR)/$(BIN_NAME)-windows-386.exe

temp=$(subst -, ,$@)
os=$(word 2, $(temp))
arch=$(subst .exe,,$(word 3, $(temp)))

.get-deps: *.go
	go get -t -d -v ./...
	touch .get-deps

all: test build

build: $(BINS)

clean:
	rm -rf .get-deps $(BIN_DIR)

test: .get-deps *.go
	go test -v -cover ./...

$(BINS): .get-deps main.go
	GOOS=$(os) GOARCH=$(arch) CGO_ENABLED=0 go build -o '$@' main.go

fmt: *.go
	go fmt *.go
