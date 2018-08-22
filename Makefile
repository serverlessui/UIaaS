GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
APP_NAME=serverlessui
SRC_LOCATION=.
BIN_OUTPUT=release
MAJOR_VERSION=0
MINOR_VERSION=0

default: clean dependencies  test build

build: serverless-ui

serverless-ui: $(GOFILES)	
	./build.sh $(APP_NAME) $(SRC_LOCATION) $(BIN_OUTPUT)

dependencies: 
	@go get github.com/tools/godep
	@dep ensure

test: test-all

test-all:
	@go test -v -cover ./$(SRC_LOCATION)/...

test-min:
	@go test ./...

clean:
	rm -rf vendor
	

publish-release:
	@go get github.com/aktau/github-release
	cd release && ./release.sh "v$(MAJOR_VERSION).$(MINOR_VERSION).$(VERSION)" $(APP_NAME)
