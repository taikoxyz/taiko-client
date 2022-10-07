
FROM golang:1.18-alpine as builder

# Github personal token to access the private repoitories.
ARG GITHUB_PERSONAL_TOKEN=""
RUN apk add --no-cache gcc musl-dev linux-headers git openssh make \
  && git config --global url.https://${GITHUB_PERSONAL_TOKEN}@github.com/taikochain.insteadOf https://github.com/taikochain

WORKDIR /client-mono
COPY . .
RUN make build

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /client-mono/bin/driver /usr/local/bin/
COPY --from=builder /client-mono/bin/proposer /usr/local/bin/
COPY --from=builder /client-mono/bin/prover /usr/local/bin/
COPY --from=builder /client-mono/script/start.sh /usr/local/bin/

ENTRYPOINT ["start.sh"]
