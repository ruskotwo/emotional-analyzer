# Building
FROM golang:1.22-alpine3.20 as builder

ADD . /go/src/github.com/ruskotwo/emotional-analyzer/tests/telegram-bot
WORKDIR /go/src/github.com/ruskotwo/emotional-analyzer/tests/telegram-bot

#RUN go mod download
RUN go build -o /go/bin/emotional-analyzer/tests/telegram-bot ./*.go

#Running
FROM alpine:3.19

COPY --from=builder /go/bin/emotional-analyzer/tests/telegram-bot /usr/local/bin/emotional-analyzer/tests/telegram-bot

ENTRYPOINT ["/usr/local/bin/emotional-analyzer/tests/telegram-bot"]