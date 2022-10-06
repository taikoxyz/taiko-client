build:
	@go build -o bin/proposer proposer/cmd/main.go \
	&& go build -o bin/prover prover/cmd/main.go

clean:
	@rm -rf bin/*

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
