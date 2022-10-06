build:
	@go build -o ./build/client-mono ./cmd/main.go

lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2 \
	&& golangci-lint run

test:
	@go test -v ./...

.PHONY: build \
				lint \
				test
