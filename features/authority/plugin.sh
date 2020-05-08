#!/bin/sh

VERSION=$1

if [ "$VERSION" = "" ]; then
VERSION=1.0
fi

NAME=authority

sed "s/github.com\/hailongz\/golang\/features\/$NAME\/$NAME/github.com\/hailongz\/golang\/features\/$NAME\/$NAME-${VERSION}/g" in.go > in-$VERSION.go

rm -rf $NAME-$VERSION
cp -r $NAME $NAME-$VERSION

rm -rf $NAME-$VERSION.so

docker run --rm -v `pwd`:/main:rw -v $GOPATH:/go:rw hailongz/kk-gobuild:latest go build  -buildmode=plugin -o $NAME-$VERSION.so in-$VERSION.go

rm -rf in-$VERSION.go
rm -rf $NAME-$VERSION
