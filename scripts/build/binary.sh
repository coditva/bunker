#!/bin/sh

mkdir -p build/
go build -o build/"$TARGET" "$TARGET.go"
