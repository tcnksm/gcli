DEBUG_FLAG = $(if $(DEBUG),-debug)

deps:
	go get github.com/jteeuwen/go-bindata/...
	go get -d -t ./...

test: deps
	go test -v ./...

install: deps
	go-bindata $(DEBUG_FLAG) -o templates.go templates
	go install
