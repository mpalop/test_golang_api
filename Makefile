.PHONY: install build run run-with-samples run-doc

# Go parameters
PACKAGE_PATH=/src/github.com/mpalop/test_golang_api
GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GODOC=godoc -http=:6060
BINARY_NAME=test_golang_api

install:
	cd $GOPATH
	$(GOGET) github.com/mpalop/sample_go_api

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

run:
	./$(BINARY_NAME)

run-with-samples:
	./$(BINARY_NAME) &
	curl -X POST -H "Content-Type: application/json" -d @$(PACKAGE_PATH)/test/fixtures/sample1.json http://localhost:8000/order
	curl -X POST -H "Content-Type: application/json" -d @$(PACKAGE_PATH)/test/fixtures/sample2.json http://localhost:8000/order

run-doc:
	$(GODOC) -http=:6060

