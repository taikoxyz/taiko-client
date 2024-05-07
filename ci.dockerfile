FROM ubuntu:latest

RUN apt update && apt install ca-certificates libstdc++6

COPY taiko-client /usr/local/bin/
RUN chmod +x /usr/local/bin/taiko-client

EXPOSE 6060

ENTRYPOINT ["taiko-client"]
