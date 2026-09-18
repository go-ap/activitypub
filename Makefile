SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

GO ?= go
TEST := $(GO) test
TEST_FLAGS ?= -v
GO111MODULE = on
PROJECT_NAME := $(shell basename $(PWD))

.PHONY: test coverage clean download

download: go.sum

go.sum: go.mod
	$(GO) mod tidy

test: go.sum
	$(TEST) $(TEST_FLAGS) -coverpkg github.com/go-ap/activitypub -cover -json ./... | go tool tparse

coverage: go.sum clean
	mkdir ./.coverage
	$(TEST) $(TEST_FLAGS) -coverpkg github.com/go-ap/activitypub -covermode=count -args -test.gocoverdir="$(PWD)/.coverage" ./... > /dev/null
	$(GO) tool covdata percent -i=./.coverage/ -o $(PROJECT_NAME).coverprofile

clean:
	@$(RM) -v *.coverprofile
	@$(RM) -r ./.coverage
