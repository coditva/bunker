#!/bin/sh

mkdir -p build/
export GO111MODULE=on
go build -o build/"$TARGET" "cmd/$TARGET/main.go"
