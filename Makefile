SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

GO ?= go
TEST := $(GO) test
TEST_FLAGS ?= -v
TEST_TARGET ?= .
GO111MODULE = on
PROJECT_NAME := $(shell basename $(PWD))

.PHONY: test coverage clean download

download: go.sum

go.sum:
	$(GO) mod tidy

test: go.sum
	$(TEST) $(TEST_FLAGS) ./tests
	$(TEST) $(TEST_FLAGS) $(TEST_TARGET)

coverage: go.sum clean
	@mkdir ./_coverage
	$(TEST) $(TEST_FLAGS) -covermode=count -args -test.gocoverdir="$(PWD)/_coverage" ./tests > /dev/null
	$(TEST) $(TEST_FLAGS) -covermode=count -args -test.gocoverdir="$(PWD)/_coverage" $(TEST_TARGET) > /dev/null
	$(GO) tool covdata percent -i=./_coverage/ -o $(PROJECT_NAME).coverprofile

clean:
	@$(RM) -v *.coverprofile
	@$(RM) -r ./_coverage
