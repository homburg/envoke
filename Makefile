.PHONY: install test

.SUFFIXES:

# Default makefile target for travis-ci.org
test:
	# go get -t >= go1.2
	go get -d -v ./... & go test -v ./...

install:
	go install .

