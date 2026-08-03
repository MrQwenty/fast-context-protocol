.PHONY: fmt vet test run build privacy-demo

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
	go build -o bin/fcpconform ./cmd/fcpconform
	go build -o bin/fcpprivacy ./cmd/fcpprivacy

privacy-demo:
	go run ./cmd/fcpprivacy -input examples/privacy/sample.txt -output /tmp/fcp-sanitized.txt -report /tmp/fcp-privacy.json -custom-terms examples/privacy/custom-terms.txt -mode anonymize
	! grep -q 'mario.rossi@example.com' /tmp/fcp-sanitized.txt
	grep -q '"passed": true' /tmp/fcp-privacy.json
