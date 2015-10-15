COMMIT = $$(git describe --always)
DEBUG_FLAG = $(if $(DEBUG),-debug)

updatedeps:
	@echo "====> Install & Update depedencies..."
	go get -v -u github.com/jteeuwen/go-bindata/...
	go get -v -d -u -t ./...

deps:
	@echo "====> Install depedencies..."
	go get -v github.com/jteeuwen/go-bindata/...
	go get -v -d -t ./...

devdeps:
	@echo "====> Install depedencies for development..."
	go get -v golang.org/x/tools/cmd/vet	
	go get -v github.com/golang/lint/golint

generate: deps
	@go generate ./...

build: generate
	@echo "====> Build gcli in ./bin "
	go build -ldflags "-X main.GitCommit=\"$(COMMIT)\"" -o bin/gcli

install: generate
	@echo "====> Install gcli in $(GOPATH)/bin ..."
	@go install -ldflags "-X main.GitCommit=\"$(COMMIT)\""

test: build devdeps
	@echo "====> Run test"
	@sh -c "$(CURDIR)/test.sh"

test-race: generate devdeps
	go test -race ./...

# Run test inside docker container
test-docker:
	@sh -c "$(CURDIR)/test-docker.sh"

# Run functional test 
test-functional: build devdeps
	@echo "====> Run functional test"
	cd tests; go test -v ./...

godoc: build
	@echo "====> Generate doc.go"
	@rm doc.go
	@./bin/gcli -godoc
