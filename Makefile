COMMIT = $$(git describe --always)
DEBUG_FLAG = $(if $(DEBUG),-debug)

deps:
	go get -v golang.org/x/tools/cmd/vet	
	go get -v github.com/golang/lint/golint
	go get -v github.com/jteeuwen/go-bindata/...
	go get -v -d -t ./...

build: deps
	cd skeleton; go-bindata -pkg="skeleton" resource/...
	go build -ldflags "-X main.GitCommit=\"$(COMMIT)\"" -o bin/gcli

install: deps
	cd skeleton; go-bindata -pkg="skeleton" resource/...
	go install -ldflags "-X main.GitCommit=\"$(COMMIT)\""

test: build
	./test.sh

tests: build
	cd tests; go test -v ./...

test-simple: build
	go test -v ./...

test-docker:
	./test-docker.sh
