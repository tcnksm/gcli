#!/bin/bash
set -e

ORG_PATH="github.com/tcnksm"
REPO_PATH="${ORG_PATH}/cli-init"

docker run \
       -v $PWD:/gopath/src/${REPO_PATH} \
       -w /gopath/src/${REPO_PATH} \
       google/golang:1.4 /bin/bash -c "make tests"
