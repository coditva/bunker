#!/bin/sh

mkdir -p build/
go build -o build/"$TARGET" "cmd/$TARGET/main.go"
