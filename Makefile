.PHONY: fmt vet test run build

fmt:
	gofmt -w $$(find . -name '*.go' -type f)

vet:
	go vet ./...

test:
	go test -race ./...

run:
	go run ./cmd/fcpd -listen :8080 -catalog examples/basic-provider/context.json

build:
	mkdir -p bin
	go build -o bin/fcpd ./cmd/fcpd
	go build -o bin/fcpctl ./cmd/fcpctl
