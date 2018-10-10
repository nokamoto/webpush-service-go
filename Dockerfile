FROM golang:1.11.0-alpine3.8 AS build

RUN apk update && apk add git

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/nokamoto/webpush-service-go

COPY Gopkg.lock .
COPY Gopkg.toml .
COPY grpc grpc
COPY *.go ./

RUN dep ensure -vendor-only=true

RUN go install .

FROM alpine:3.8

RUN apk update && apk add --no-cache ca-certificates

COPY --from=build /go/bin/webpush-service-go /usr/local/bin/webpush-service-go

ENTRYPOINT [ "webpush-service-go" ]
