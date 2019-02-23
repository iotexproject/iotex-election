FROM golang:1.11.5-stretch

WORKDIR $GOPATH/src/github.com/iotexproject/iotex-election/

RUN apt-get install -y --no-install-recommends make

COPY . .

ARG SKIP_DEP=false

RUN if [ "$SKIP_DEP" != true ] ; \
    then \
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
        dep ensure --vendor-only; \
    fi

RUN rm -rf ./bin/server && \
    rm -rf election.db && \
    go build -o ./bin/server -v ./server && \
    cp $GOPATH/src/github.com/iotexproject/iotex-election/bin/server /usr/local/bin/iotex-server  && \
    mkdir -p /etc/iotex/ && \
    cp server.yaml /etc/iotex/server.yaml && \
    rm -rf $GOPATH/src/github.com/iotexproject/iotex-election/

CMD [ "iotex-server", "-config=/etc/iotex/server.yaml"]