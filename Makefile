DEBUG_FLAG = $(if $(DEBUG),-debug)

deps:
	go get -v -u github.com/jteeuwen/go-bindata/...
	go get -v -d -t ./...

test: deps
	go test -v ./...

build: deps
	cd skeleton; go-bindata -pkg="skeleton" resource/...
	go build -o bin/gcli

install: deps
	cd skeleton; go-bindata -pkg="skeleton" resource/...
	go install

tests: build
	cd tests; go test -v ./...

