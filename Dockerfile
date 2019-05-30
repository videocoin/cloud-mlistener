FROM alpine:3.7

COPY bin/mlistener /opt/videocoin/bin/mlistener

CMD ["/opt/videocoin/bin/mlistener"]
