FROM instrumentisto/dep

ENV GOBIN /go/bin
ENV PEER_PORT ":8080"

RUN mkdir /go/src/simple-poa-blockchain-peer
COPY . /go/src/simple-poa-blockchain-peer/
WORKDIR /go/src/simple-poa-blockchain-peer/

RUN dep ensure && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/src/simple-poa-blockchain-peer/main .

EXPOSE 8080

ENTRYPOINT ["/go/src/simple-poa-blockchain-peer/main"]

