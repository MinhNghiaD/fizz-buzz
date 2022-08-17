-include .env

PROJECTNAME := $(shell basename "$(PWD)")
SHELL:=/bin/bash
BIN := "./bin"
DOCKERFOLDER := ./docker
VERSION := $(shell git describe --tags --always)

GOLINTPATH := `(go env GOPATH)` 

.PHONY: vendor
vendor:
	@echo "  > Update modules" 
	@go mod tidy && go mod vendor

.PHONY: build
build:
	@-$(MAKE) clean
	@echo "  >  Building binary..."
	@CGO_ENABLED=0 go build -mod=vendor -o $(BIN)/ ./...

.PHONY: docker
docker:	
	@docker build . -f $(DOCKERFOLDER)/Dockerfile -t $(PROJECTNAME):$(VERSION)

.PHONY: test
test:	
	@go test -race ./pkg/./... 

.PHONY: clean
clean:
	@echo "  >  Cleaning build cache"
	@go clean

.PHONY: lint
lint:
	export GOPATH=$(GOLINTPATH) && golangci-lint run $(CMD_RECURSIVE) --timeout=2m