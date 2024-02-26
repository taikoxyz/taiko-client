FROM gcr.io/prysmaticlabs/prysm/cmd/prysmctl:latest AS prysmctl

FROM gcr.io/prysmaticlabs/prysm/beacon-chain:stable AS beacon

FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install bc -y

COPY --from=prysmctl /prysmctl /usr/local/bin
COPY --from=beacon /beacon-chain /usr/local/bin

COPY genesis.json /
COPY config.yml /
COPY jwtsecret /
COPY start.sh /

ENTRYPOINT ["/start.sh"]
