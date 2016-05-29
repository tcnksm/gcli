#!/bin/bash
# This script runs test (`gofmt`, `golint`, `go vet` & `go test`)
# If `gofmt` & `golint` has output (means something wrong),
# it will exit with non-zero status

TARGET=$(find . -name "*.go" | grep -v "bindata.go" | grep -v "doc.go")
echo -e "----> Run gofmt"
FMT_RES=$(gofmt -l ${TARGET})
if [ -n "${FMT_RES}" ]; then
    echo -e "gofmt failed: \n${FMT_RES}"
    exit 255
fi

echo -e "----> Run go vet"
go list -f '{{.Dir}}' ./... | xargs go tool vet
if [ $? -ne 0 ]; then
    echo -e "go vet failed"
    exit 255
fi

# TODO, better way to exclude some lint warning.
echo -e "----> Run golint"
LINT_RES=$(golint ./... | \
                  grep -v "bindata.go" | \
                  grep -v "doc.go" | \
                  grep -v "type name will be used as command.CommandFlag by other packages" | \
                  grep -v "Framework_go_cmd" | \
                  grep -v "Framework_urfave_cli" | \
                  grep -v "Framework_mitchellh_cli" | \
                  grep -v "Framework_flag" \
        )
if [ -n "${LINT_RES}" ]; then
     echo -e "golint failed: \n${LINT_RES}"
fi

echo -e "----> Run go test"
go test -v ./...
