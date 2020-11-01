#!/bin/bash
set -e

function build() {
  mkdir -p dist
  local os=$1
  local arch=$2
  local target=ssh443client
  local ext=''

  if [[ $os = windows ]]; then
    ext=.exe
  fi

  GOOS=$os GOARCH=$arch \
    go build \
    -o dist/${os}_${arch}_${target}${ext} \
    $target.go
}

build darwin 386
build darwin amd64
build dragonfly amd64
build freebsd 386
build freebsd amd64
build freebsd arm
build linux 386
build linux amd64
build linux arm64
build linux arm
build linux mips64
build linux mips64le
build linux mips
build linux mipsle
build linux ppc64
build linux ppc64le
build netbsd 386
build netbsd amd64
build netbsd arm
build openbsd 386
build openbsd amd64
build openbsd arm
build plan9 386
build plan9 amd64
build solaris amd64
build windows 386
build windows amd64
