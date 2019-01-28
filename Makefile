TEST := go test
TEST_FLAGS ?= -v
TEST_TARGET ?= ./...
GO111MODULE=on

test:
	$(TEST) $(TEST_FLAGS) $(TEST_TARGET)

coverprofile: TEST_TARGET := .
coverprofile: TEST_FLAGS += -covermode=count -coverprofile=$(TEST_TARGET).coverprofile
coverprofile: test

clean:
	$(RM) -v *.coverprofile

