.PHONY: install test

.SUFFIXES:

# Default makefile target for travis-ci.org
test:
	go get -t -u && go test -v ./...

install:
	go install .

