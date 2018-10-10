#!/bin/sh

set -ex

PUSH=false
VERSION=0.0.0
TAG=nokamoto13/webpush-service-go:$VERSION

for arg in "$@"
do
    if [ "$arg" = "--push" ]
    then
        PUSH=true
    fi
done

docker build -t $TAG .

if "$PUSH"
then
    docker push $TAG
fi
