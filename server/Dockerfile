FROM golang:1.14 AS builder

WORKDIR /go/src/github.com/standielpls/articulate
COPY . .
RUN export VERSION=$(cat VERSION); \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o $GOPATH/bin/articulate cmd/main.go

FROM alpine:3.11

RUN apk add --update --no-cache ca-certificates postgresql-client
ENV PATH="${PATH}":/go/bin
COPY --from=builder /go/bin /go/bin

ENV ADDR 8080

CMD ["articulate"]