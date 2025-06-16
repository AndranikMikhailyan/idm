.DEFAULT := build

.PHONY: fmt vet build

test:
	go test ./inner...

fmt: test
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build