SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

GO ?= go
TEST := $(GO) test
TEST_FLAGS ?= -v
TEST_TARGET ?= ./...
GO111MODULE = on
PROJECT := $(shell basename $(PWD))

ifeq ($(origin GOCOVERDIR),undefined)
GOCOVERDIR := $(shell pwd)/_coverage
endif

.PHONY: test coverage clean download

download: go.sum

go.sum: go.mod
	$(GO) mod tidy

test: go.sum
	$(TEST) $(TEST_FLAGS) -cover -json $(TEST_TARGET) | go tool tparse

$(GOCOVERDIR):
	mkdir -p $(GOCOVERDIR)

coverage: TEST_FLAGS += -coverpkg github.com/go-ap/activitypub -covermode=count -args -test.gocoverdir=$(GOCOVERDIR)
coverage: $(GOCOVERDIR)
	$(TEST) $(TEST_FLAGS)  $(TEST_TARGET) > /dev/null
	$(GO) tool covdata percent -i=$(GOCOVERDIR) -o $(PROJECT).coverprofile
	$(MAKE) clean

clean:
	@$(RM) -v *.coverprofile
	$(RM) -r $(GOCOVERDIR)
