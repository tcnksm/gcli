#!/bin/bash
# This script runs test (`gofmt`, `golint`, `go vet` & `go test`)
# If `gofmt` & `golint` has output (means something wrong),
# it will exit with non-zero status

TARGET=$(find . -name "*.go" | grep -v "bindata.go" | grep -v "doc.go" | grep vendor)
echo -e "----> Run gofmt"
FMT_RES=$(gofmt -l ${TARGET})
if [ -n "${FMT_RES}" ]; then
    echo -e "gofmt failed: \n${FMT_RES}"
    exit 255
fi

echo -e "----> Run go vet"
go vet $(go list ./... | grep -v /vendor/)
if [ $? -ne 0 ]; then
    echo -e "go vet failed"
    exit 255
fi

# TODO, better way to exclude some lint warning.
echo -e "----> Run golint"
go list ./... | grep -v /vendor/ | xargs -L1 golint | grep -v bindata.go | grep -v doc.go | grep -v "type name will be used as command.CommandFlag by other packages"

echo -e "----> Run go test"
go test -v $(go list ./... | grep -v /vendor/)
