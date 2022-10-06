build:
	@go build -o ./build/client-mono ./cmd/main.go

lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2 \
	&& golangci-lint run

test:
	@TAIKO_MONO_DIR=${TAIKO_MONO_DIR} \
	COMPILE_PROTOCOL=${COMPILE_PROTOCOL} \
		./integration_tests/entrypoint.sh

.PHONY: build \
				lint \
				test
