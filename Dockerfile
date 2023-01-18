FROM golang:1.18-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git make

WORKDIR /taiko-client
COPY . .
RUN make build

RUN git clone --depth 1 --branch feature/root-circuit https://github.com/smtmfft/zkevm-circuits.git /zkevm-circuits

RUN cd /zkevm-circuits && ./build_pi_integration.sh && \
  chmod +x ./pi_circuit_integration && \
  cp ./pi_circuit_integration /usr/local/bin/pi_circuit_integration

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /taiko-client/bin/taiko-client /usr/local/bin/
COPY --from=builder /zkevm-circuits/pi_circuit_integration /usr/local/bin/

EXPOSE 6060

ENTRYPOINT ["taiko-client"]
