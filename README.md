# webpush-service-go
[![CircleCI](https://circleci.com/gh/nokamoto/webpush-service-go.svg?style=svg)](https://circleci.com/gh/nokamoto/webpush-service-go)

## Build
```bash
protoc --go_out=plugins=grpc:grpc -I webpush-protobuf webpush-protobuf/webpush/protobuf/*.proto
```
