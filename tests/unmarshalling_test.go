package tests

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"

	a "activitypub"
	j "jsonld"
)

const dir = "./mocks"

var stopOnFailure = false

type testPair struct {
	path     string
	expected bool
	blank    a.ObjectOrLink
	result   a.ObjectOrLink
}

type tests map[string]testPair

var allTests tests = tests{
	"nil": testPair{
		path:     "./mocks/empty.json",
		expected: true,
		blank:    a.Object{},
		result:   a.Object{},
	},
}

func getFileContents(path string) []byte {
	f, _ := os.Open(path)

	data := make([]byte, 512)
	io.ReadFull(f, data)
	data = bytes.Trim(data, "\x00")

	return data
}

func Test_ActivityPubUnmarshall(t *testing.T) {
	var err error

	var f = t.Errorf
	if stopOnFailure {
		f = t.Fatalf
	}

	if len(allTests) == 0 {
		t.Skip("No tests found")
	}

	for _, pair := range allTests {
		fmt.Printf("===          %s\n", pair.path)

		data := getFileContents(pair.path)

		object := pair.blank

		err = j.Unmarshal(data, object)
		if err != nil {
			f("Error: %s", err)
			continue
		}
		expLbl := ""
		if !pair.expected {
			expLbl = "not be "
		}
		if pair.expected != reflect.DeepEqual(object, pair.result) {
			f("\n%#v\n should %sequal to\n%#v", object, expLbl, pair.result)
			continue
		}
		if err == nil {
			fmt.Print("---          PASS\n")
		}
	}
}
