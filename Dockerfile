FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/Swapica/aggregator-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/aggregator-svc /go/src/github.com/Swapica/aggregator-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/aggregator-svc /usr/local/bin/aggregator-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["aggregator-svc"]
