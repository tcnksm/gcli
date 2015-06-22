#!/bin/bash
# This script runs test (`gofmt`, `golint`, `go vet` & `go test`)
# If `gofmt` & `golint` has output (means something wrong),
# it will exit with non-zero status

TARGET=$(find . -name "*.go" | grep -v "bindata.go")
FMT_RES=$(gofmt -l ${TARGET})
if [ -n "${FMT_RES}" ]; then
    echo -e "gofmt failed: \n${FMT_RES}"
    exit 255
fi

go vet ./...
if [ $? -ne 0 ]; then
    echo -e "go vet failed"
    exit 255
fi

LINT_RES=$(golint ./...)
if [ -n "${LINT_RES}" ]; then
     echo -e "golint failed: \n${LINT_RES}"
     exit 255
fi

go test -v ./...
