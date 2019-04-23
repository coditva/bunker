#!/bin/sh

export GO111MODULE=on
go test "cmd/$TARGET/main.go"
