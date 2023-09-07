files := $(shell find . -name "*.go" | grep -v vendor)

bootstrap:
	go install -v golang.org/x/lint/golint@latest
	go install -v golang.org/x/tools/...@latest
	go install -v honnef.co/go/tools/cmd/staticcheck@latest

clean:
	goimports -w $(files)

test: clean
	go test

lint:
	golint -set_exit_status
	golint -set_exit_status example
	staticcheck github.com/zerodivisi0n/techan
	staticcheck github.com/zerodivisi0n/techan/example

bench: clean
	go test -bench .

commit: test
	git commit

release: clean test lint
	./scripts/release.sh

test-with-coverage:
	go test -race -cover -covermode=atomic -coverprofile=coverage.txt github.com/zerodivisi0n/techan

view-coverage: test-with-coverage
	go tool cover -html coverage.txt
