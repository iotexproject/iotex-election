FROM golang:1.14-alpine as build
RUN apk add --no-cache make gcc musl-dev linux-headers git
ENV GO111MODULE=on

WORKDIR apps/iotex-election

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./bin/server -v . && \
    cp ./bin/server /usr/local/bin/iotex-server  && \
    mkdir -p /etc/iotex/ && \
    cp server.yaml /etc/iotex/server.yaml

FROM alpine:latest

RUN apk add --no-cache ca-certificates

RUN mkdir -p /etc/iotex/
COPY --from=build /etc/iotex/server.yaml /etc/iotex
COPY --from=build /usr/local/bin/iotex-server /usr/local/bin

CMD [ "iotex-server", "-config=/etc/iotex/server.yaml"]
