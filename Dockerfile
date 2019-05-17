FROM golang:1.12.5-stretch

WORKDIR apps/iotex-election

RUN apt-get install -y --no-install-recommends make

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN rm -rf ./bin/server && \
    rm -rf election.db && \
    go build -o ./bin/server -v . && \
    cp ./bin/server /usr/local/bin/iotex-server  && \
    mkdir -p /etc/iotex/ && \
    cp server.yaml /etc/iotex/server.yaml && \
    rm -rf apps/iotex-election

CMD [ "iotex-server", "-config=/etc/iotex/server.yaml"]
