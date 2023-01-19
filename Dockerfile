FROM golang:1.18

RUN apt-get update && apt-get install -y git musl-dev curl build-essential make ca-certificates && \
  curl https://sh.rustup.rs -sSf | bash -s -- -y

ENV PATH="/root/.cargo/bin:${PATH}"

WORKDIR /taiko-client
COPY . .
RUN make build

RUN git clone --branch feature/root-circuit https://github.com/smtmfft/zkevm-circuits.git /zkevm-circuits

WORKDIR /zkevm-circuits
RUN git reset bd22fc2 && \
  ./build_pi_integration.sh && \
  chmod +x ./pi_circuit_integration && \
  cp /zkevm-circuits/pi_circuit_integration /usr/local/bin/ && \
  cp /taiko-client/bin/taiko-client /usr/local/bin/

EXPOSE 6060

ENTRYPOINT ["taiko-client"]
