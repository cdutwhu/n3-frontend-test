#!/bin/bash

set -e
shopt -s extglob

rm -rf ./build

go get

GOARCH=amd64
LDFLAGS="-s -w"
OUT=server

# OUTPATH=./build/win64/
# mkdir -p $OUTPATH
# GOOS="windows" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT.exe
# mv $OUT.exe $OUTPATH
# cp ./config/*.toml $OUTPATH
# cp -r ./www $OUTPATH

# OUTPATH=./build/mac/
# mkdir -p $OUTPATH
# GOOS="darwin" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
# mv $OUT $OUTPATH
# cp ./config/*.toml $OUTPATH
# cp -r ./www $OUTPATH

OUTPATH=./build/linux64/
mkdir -p $OUTPATH
GOOS="linux" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
mv $OUT $OUTPATH
cp ./config/*.toml $OUTPATH
cp -r ./www $OUTPATH

# GOARCH=arm
# OUTPATH=./build/linuxarm/
# mkdir -p $OUTPATH
# GOOS="linux" GOARCH="$GOARCH" GOARM=7 go build -ldflags="$LDFLAGS" -o $OUT
# mv $OUT $OUTPATH
# cp ./config/*.toml $OUTPATH
# cp -r ./www $OUTPATH
