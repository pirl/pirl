# Build Geth in a stock Go builder container
FROM golang:1.9-alpine as builder
##FROM ubuntu
RUN apk add --update bash && rm -rf /var/cache/apk/*
RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /pirl
RUN cd /pirl && make pirl

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /pirl/build/bin/pirl /usr/local/bin/

ADD start.sh /root/start.sh

EXPOSE 6588 6589 30303 30303/udp 30304/udp
ENTRYPOINT /root/start.sh
