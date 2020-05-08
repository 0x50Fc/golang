#!/bin/sh

VERSION=$1

if [ "$VERSION" = "" ]; then
VERSION=1.0
fi

NAME=wx

sed "s/github.com\/hailongz\/golang\/features\/$NAME\/$NAME/github.com\/hailongz\/golang\/features\/$NAME\/$NAME-${VERSION}/g" in.go > in-$VERSION.go

rm -rf $NAME-$VERSION
ln -s `pwd`/$NAME  $NAME-$VERSION
rm -rf $NAME-$VERSION-local.so

go build -buildmode=plugin -o $NAME-$VERSION-local.so in-$VERSION.go

rm -rf in-$VERSION.go
rm -rf $NAME-$VERSION

rm -rf ../../xs/app/iid/$NAME-$VERSION-local.so
cp $NAME-$VERSION-local.so ../../xs/app/iid/$NAME-$VERSION-local.so
