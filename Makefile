DEBUG_FLAG = $(if $(DEBUG),-debug)

deps:
	go get -v github.com/jteeuwen/go-bindata/...
	go get -v -d -t ./...

build: deps
	cd skeleton; go-bindata -pkg="skeleton" resource/...
	go build -o bin/gcli

install: deps
	cd skeleton; go-bindata -pkg="skeleton" resource/...
	go install

test: build
	go test ./...

tests: build
	cd tests; go test -v ./...

