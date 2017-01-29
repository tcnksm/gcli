VERSION = $(shell grep 'Version string' version.go | sed -E 's/.*"(.+)"$$/\1/')
COMMIT = $(shell git describe --always)
PACKAGES = $(shell go list ./... | grep -v '/vendor/')
EXTERNAL_TOOLS = github.com/jteeuwen/go-bindata

default: test

# install external tools for this project
bootstrap:
	@echo "====> Install & Update depedencies..."
	@for tool in $(EXTERNAL_TOOLS) ; do \
		echo "Installing $$tool" ; \
    	go get $$tool; \
	done

generate: 
	@go generate ./...

godoc: build
	@echo "====> Generate doc.go"
	@rm doc.go
	@./bin/gcli -godoc

build: generate
	@echo "====> Build gcli in ./bin "
	go build -ldflags "-X main.GitCommit=\"$(COMMIT)\"" -o bin/gcli

install: generate
	@echo "====> Install gcli in $(GOPATH)/bin ..."
	@go install -ldflags "-X main.GitCommit=\"$(COMMIT)\""

.PHONY: bootstrap generate godoc build install

test-all: vet lint test test-race

test: build
	go test -v -parallel=4 ${PACKAGES}

test-race:
	go test -v -race ${PACKAGES}

# Run functional test 
test-functional: build devdeps
	@echo "====> Run functional test"
	cd tests; go test -v ./...

vet:
	go vet ${PACKAGES}

lint:
	@go get github.com/golang/lint/golint
	go list ./... | grep -v vendor | xargs -n1 golint |\
		grep -v godoc.go |\
		grep -v bindata.go |\		

cover:
	@go get golang.org/x/tools/cmd/cover		
	go test -coverprofile=cover.out
	go tool cover -html cover.out
	rm cover.out

.PHONY: test test-race test-functional test-all vet lint cover  


