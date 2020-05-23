#!/bin/bash

set -e
shopt -s extglob

rm -rf ./build/*
mkdir -p ./build/Linux64 ./build/Win64 ./build/Mac ./build/LinuxArm

go get

GOARCH=amd64
LDFLAGS="-s -w"
OUT=n3-test

OUTPATH=./build/Win64/
GOOS="windows" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT.exe
mv $OUT.exe $OUTPATH
cp ./config/config.toml $OUTPATH

OUTPATH=./build/Mac/
GOOS="darwin" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
mv $OUT $OUTPATH
cp ./config/config.toml $OUTPATH

OUTPATH=./build/Linux64/
GOOS="linux" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
mv $OUT $OUTPATH
cp ./config/config.toml $OUTPATH

GOARCH=arm
OUTPATH=./build/LinuxArm/
GOOS="linux" GOARCH="$GOARCH" GOARM=7 go build -ldflags="$LDFLAGS" -o $OUT
mv $OUT $OUTPATH
cp ./config/config.toml $OUTPATH