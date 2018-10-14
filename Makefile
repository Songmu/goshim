
test: deps
	go test

deps:
	go get -d -v -t ./...
	go get golang.org/x/lint/golint
	go get github.com/mattn/goveralls

lint: deps
	go vet ./...
	golint -set_exit_status ./...

cover: deps
	go get github.com/axw/gocov/gocov
	goveralls

.PHONY: test deps lint cover
