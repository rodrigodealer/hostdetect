FROM golang:1.11-alpine as builder
COPY . /go/src/brfutebol/hostdetect
ENV GO111MODULE=on
WORKDIR /go/src/brfutebol/hostdetect
RUN apk -U add git build-base && \
    mkdir -p /build && \
    go build  -ldflags '-extldflags "-static"' -o /build/hostdetect

FROM alpine:3.7
RUN apk -U add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /opt
COPY --from=builder /build/hostdetect .
CMD ["./opt/hostdetect"]