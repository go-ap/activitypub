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
	"object_with_url": testPair{
		path:     "./mocks/object_with_url.json",
		expected: true,
		blank:    &a.Object{},
		result: &a.Object{
			URL: a.URI("http://littr.git/api/accounts/system"),
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
	//	expected: false,
	//	blank:    &a.Activity{},
	//	result: &a.Activity{
	//		Type:    a.ActivityType,
	//		Summary: a.NaturalLanguageValue{a.LangRef("-"): "Sally did something to a note"},
	//		Actor: a.Actor(a.Person{
	//			Type: a.PersonType,
	//			Name: a.NaturalLanguageValue{
	//				a.LangRef("-"): "Sally",
	//			},
	//		}),
	//		Object: a.Object{
	//			Type: a.NoteType,
	//			Name: a.NaturalLanguageValue{
	//				a.LangRef("-"): "A Note",
	//			},
	//		},
	//	},
	//},
	"person_with_outbox": testPair{
		path:     "./mocks/person_with_outbox.json",
		expected: true,
		blank:    &a.Person{},
		result: &a.Person{
			ID:   a.ObjectID("http://example.com/accounts/ana"),
			Type: a.PersonType,
			Name: a.NaturalLanguageValue{
				a.LangRef("-"): "ana",
			},
			PreferredUsername: a.NaturalLanguageValue{
				a.LangRef("-"): "Ana",
			},
			URL: a.URI("http://example.com/accounts/ana"),
			Outbox: a.OutboxStream{
				ID:   a.ObjectID("http://example.com/accounts/ana/outbox"),
				Type: a.OrderedCollectionType,
				URL:  a.URI("http://example.com/outbox"),
			},
		},
	},
}

func getFileContents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data := make([]byte, 512)
	io.ReadFull(f, data)
	data = bytes.Trim(data, "\x00")

	return data, nil
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
		var data []byte
		data, err = getFileContents(pair.path)
		if err != nil {
			f("Error: %s for %s", err, pair.path)
			continue
		}
		object := pair.blank

		err = j.Unmarshal(data, object)
		if err != nil {
			f("Error: %s for %s", err, data)
			continue
		}
		expLbl := ""
		if !pair.expected {
			expLbl = "not be "
		}
		if pair.expected != reflect.DeepEqual(object, pair.result) {
			f("\n%#v\n should %sequal to\n%#v", object, expLbl, pair.result)
			oj, e1 := j.Marshal(object)
			if e1 != nil {
				f(e1.Error())
			}
			tj, e2 := j.Marshal(pair.result)
			if e2 != nil {
				f(e2.Error())
			}

			f("\n%s\n should %sequal to\n%s", oj, expLbl, tj)
			continue
		}
		//fmt.Printf("%#v", object)
		if err == nil {
			fmt.Printf(" --- %s: %s\n          %s\n", "PASS", k, pair.path)
		}
	}
}
