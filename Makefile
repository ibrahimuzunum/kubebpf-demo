.PHONY: build run test

build:
	mkdir -p bin
	go build -o bin/kubebpf ./cmd/kubebpf

run:
	go run ./cmd/kubebpf doctor

test:
	go test ./...
