build:
	@go build -o ./bin/client-mono ./cmd/main.go

clean:
	@rm -rf bin/client-mono

lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2 \
	&& golangci-lint run

test:
	@TAIKO_MONO_DIR=${TAIKO_MONO_DIR} \
	COMPILE_PROTOCOL=${COMPILE_PROTOCOL} \
		./integration_test/entrypoint.sh

gen_bindings:
	@TAIKO_MONO_DIR=${TAIKO_MONO_DIR} \
	TAIKO_CLIENT_DIR=${TAIKO_CLIENT_DIR} \
		./script/gen_bindings.sh

.PHONY: build \
				clean \
				lint \
				test \
				gen_bindings
