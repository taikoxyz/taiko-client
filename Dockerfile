FROM golang:1.18 as builder

RUN apt-get update && apt-get install -y git musl-dev curl build-essential make && \
  curl https://sh.rustup.rs -sSf | bash -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"

WORKDIR /taiko-client
COPY . .
RUN make build

RUN git clone --depth 1 --branch feature/root-circuit https://github.com/smtmfft/zkevm-circuits.git /zkevm-circuits

WORKDIR /zkevm-circuits
RUN ./build_pi_integration.sh && \
  chmod +x ./pi_circuit_integration && \
  cp ./pi_circuit_integration /usr/local/bin/pi_circuit_integration

FROM node:16

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /taiko-client/bin/taiko-client /usr/local/bin/
COPY --from=builder /zkevm-circuits/pi_circuit_integration /usr/local/bin/

EXPOSE 6060

ENTRYPOINT ["taiko-client"]
