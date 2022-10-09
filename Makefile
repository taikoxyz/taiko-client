build:
	@go build -o bin/taiko-client cmd/main.go

clean:
	@rm -rf bin/*

lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2 \
	&& golangci-lint run

test:
	@TAIKO_MONO_DIR=${TAIKO_MONO_DIR} \
	COMPILE_PROTOCOL=${COMPILE_PROTOCOL} \
	RUN_TESTS=true \
		./integration_test/entrypoint.sh

dev_net:
	@TAIKO_MONO_DIR=${TAIKO_MONO_DIR} \
	COMPILE_PROTOCOL=${COMPILE_PROTOCOL} \
		./integration_test/entrypoint.sh

gen_bindings:
	@TAIKO_MONO_DIR=${TAIKO_MONO_DIR} \
	TAIKO_GETH_DIR=${TAIKO_GETH_DIR} \
		./script/gen_bindings.sh

.PHONY: build \
				clean \
				lint \
				test \
				dev_net \
				gen_bindings
