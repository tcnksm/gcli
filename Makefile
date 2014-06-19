deps:
	go get -d -t ./...

install: deps
	go install
