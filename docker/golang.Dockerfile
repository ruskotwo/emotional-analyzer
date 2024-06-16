# Building
FROM golang:1.22-alpine3.20 as builder

ADD . /go/src/github.com/ruskotwo/emotional-analyzer
WORKDIR /go/src/github.com/ruskotwo/emotional-analyzer

RUN go install github.com/google/wire/cmd/wire@latest
RUN cd cmd/factory && wire ; cd ../..

#RUN go mod download
RUN go build -o /go/bin/emotional-analyzer ./cmd/main.go

#Running
FROM alpine:3.19

COPY --from=builder /go/bin/emotional-analyzer /usr/local/bin/emotional-analyzer

ENTRYPOINT ["/usr/local/bin/emotional-analyzer"]