TEST := go test
TEST_FLAGS := -v
TEST_TARGET := ./...

export GOPATH += $$GOPATH:$(shell pwd)

test:
	$(TEST) $(TEST_FLAGS) $(TEST_TARGET)

activitypub.coverprofile: TEST_TARGET := activitypub
activitypub.coverprofile: TEST_FLAGS += -covermode=count -coverprofile=$(TEST_TARGET).coverprofile
activitypub.coverprofile: go get -v -u github.com/buger/jsonparser
activitypub.coverprofile: test

activitypub.coverprofile: TEST_TARGET := jsonld
activitypub.coverprofile: TEST_FLAGS += -covermode=count -coverprofile=$(TEST_TARGET).coverprofile
activitypub.coverprofile: test

clean:
	$(RM) -v *.coverprofile

