FROM golang:1.12.4 as builder
WORKDIR /go/src/github.com/videocoin/cloud-mlistener
COPY . .
RUN make build

FROM bitnami/minideb:jessie
COPY --from=builder /go/src/github.com/videocoin/cloud-mlistener/bin/mlistener /opt/videocoin/bin/mlistener
CMD ["/opt/videocoin/bin/mlistener"]