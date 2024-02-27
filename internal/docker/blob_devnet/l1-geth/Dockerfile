FROM gcr.io/prysmaticlabs/prysm/cmd/prysmctl:latest AS prysmctl

FROM ethereum/client-go:v1.13.13 AS geth

FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install bc -y

COPY --from=prysmctl /prysmctl /usr/local/bin
COPY --from=geth /usr/local/bin/geth /usr/local/bin

WORKDIR execution

COPY start.sh .
COPY genesis.json .
COPY keyfile.json .
COPY geth_password.txt .
COPY jwtsecret .
COPY config.yml .

ENTRYPOINT ["sh", "/execution/start.sh"]
