package tests

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"

	a "github.com/mariusor/activitypub.go/activitypub"
	j "github.com/mariusor/activitypub.go/jsonld"
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

var allTests = tests{
	"empty": testPair{
		path:     "./mocks/empty.json",
		expected: true,
		blank:    &a.Object{},
		result:   &a.Object{},
	},
	"link_simple": testPair{
		path:     "./mocks/link_simple.json",
		expected: true,
		blank:    &a.Link{},
		result: &a.Link{
			Type:      a.LinkType,
			Href:      a.URI("http://example.org/abc"),
			HrefLang:  a.LangRef("en"),
			MediaType: a.MimeType("text/html"),
			Name: a.NaturalLanguageValue{
				a.LangRef("-"): "An example link",
			},
		},
	},
	"object_simple": testPair{
		path:     "./mocks/object_simple.json",
		expected: true,
		blank:    &a.Object{},
		result: &a.Object{
			Type: a.ObjectType,
			ID:   a.ObjectID("http://www.test.example/object/1"),
			Name: a.NaturalLanguageValue{
				a.LangRef("-"): "A Simple, non-specific object",
			},
		},
	},
	//"activity_simple": testPair{
	//	path:     "./mocks/activity_simple.json",
	//	expected: true,
	//	blank:    &a.Activity{},
	//	result: &a.Activity{
	//		Type:    a.ActivityType,
	//		Summary: a.NaturalLanguageValue{a.LangRef("-"): "Sally did something to a note"},
	//	},
	//},
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

	for k, pair := range allTests {
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
			fmt.Printf(" --- %s: %s\n          %s\n", "PASS", k, pair.path)
		}
	}
}
