# Building
FROM golang:1.22-alpine3.20 as builder

RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.1

ADD . /go/src/github.com/ruskotwo/emotional-analyzer
WORKDIR /go/src/github.com/ruskotwo/emotional-analyzer

#RUN go mod download
RUN go build -o /go/bin/emotional-analyzer ./cmd/main.go

#Running
FROM alpine:3.19

COPY --from=builder /go/bin/emotional-analyzer /usr/local/bin/emotional-analyzer
COPY --from=builder /go/bin/dlv /go/bin/dlv

ENTRYPOINT exec /go/bin/dlv --listen=:40001 --continue --accept-multiclient --headless=true --api-version=2 --check-go-version=false exec /usr/local/bin/emotional-analyzer