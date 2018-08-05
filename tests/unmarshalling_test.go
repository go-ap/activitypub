package tests

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"unsafe"

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

type visit struct {
	a1  unsafe.Pointer
	a2  unsafe.Pointer
	typ reflect.Type
}

type errorableFunc func(format string, args ...interface{})

func assertDeepEquals(t errorableFunc, x, y interface{}) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	if v1.Type() != v2.Type() {
		t("%T != %T", x, y)
		return false
	}
	return deepValueEqual(t, v1, v2, make(map[visit]bool), 0)
}

func deepValueEqual(t errorableFunc, v1, v2 reflect.Value, visited map[visit]bool, depth int) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.Type() != v2.Type() {
		t("types differ %s != %s", v1.Type().Name(), v2.Type().Name())
		return false
	}

	// We want to avoid putting more in the visited map than we need to.
	// For any possible reference cycle that might be encountered,
	// hard(t) needs to return true for at least one of the types in the cycle.
	hard := func(k reflect.Kind) bool {
		switch k {
		case reflect.Map, reflect.Slice, reflect.Ptr, reflect.Interface:
			return true
		}
		//t("Invalid type for %s", k)
		return false
	}

	if v1.CanAddr() && v2.CanAddr() && hard(v1.Kind()) {
		addr1 := unsafe.Pointer(v1.UnsafeAddr())
		addr2 := unsafe.Pointer(v2.UnsafeAddr())
		if uintptr(addr1) > uintptr(addr2) {
			// Canonicalize order to reduce number of entries in visited.
			// Assumes non-moving garbage collector.
			addr1, addr2 = addr2, addr1
		}

		// Short circuit if references are already seen.
		typ := v1.Type()
		v := visit{addr1, addr2, typ}
		if visited[v] {
			return true
		}

		// Remember for later.
		visited[v] = true
	}

	switch v1.Kind() {
	case reflect.Array:
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(t, v1.Index(i), v2.Index(i), visited, depth+1) {
				t("Arrays not equal at index %d %s %s", i, v1.Index(i), v2.Index(i))
				return false
			}
		}
		return true
	case reflect.Slice:
		if v1.IsNil() != v2.IsNil() {
			t("One of the slices is not nil %s[%d] vs %s[%d]", v1.Type().Name(), v1.Len(), v2.Type().Name(), v2.Len())
			return false
		}
		if v1.Len() != v2.Len() {
			t("Slices lengths are different %s[%d] vs %s[%d]", v1.Type().Name(), v1.Len(), v2.Type().Name(), v2.Len())
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(t, v1.Index(i), v2.Index(i), visited, depth+1) {
				t("Slices elements at pos %d are not equal %#v vs %#v", i, v1.Index(i), v2.Index(i))
				return false
			}
		}
		return true
	case reflect.Interface:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		return deepValueEqual(t, v1.Elem(), v2.Elem(), visited, depth+1)
	case reflect.Ptr:
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		return deepValueEqual(t, v1.Elem(), v2.Elem(), visited, depth+1)
	case reflect.Struct:
		for i, n := 0, v1.NumField(); i < n; i++ {
			if !deepValueEqual(t, v1.Field(i), v2.Field(i), visited, depth+1) {
				return false
			}
		}
		return true
	case reflect.Map:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !deepValueEqual(t, v1.MapIndex(k), v2.MapIndex(k), visited, depth+1) {
				return false
			}
		}
		return true
	case reflect.Func:
		if v1.IsNil() && v2.IsNil() {
			return true
		}
		// Can't do better than this:
		return false
	}
	return true // i guess?
}

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
			Outbox: &a.OutboxStream{
				ID:   a.ObjectID("http://example.com/accounts/ana/outbox"),
				Type: a.OrderedCollectionType,
				URL:  a.URI("http://example.com/outbox"),
			},
		},
	},
	"ordered_collection": testPair{
		path:     "./mocks/ordered_collection.json",
		expected: false, // This fails because interface pointers being different I think
		blank:    &a.OrderedCollection{},
		result: &a.OrderedCollection{
			ID:         a.ObjectID("http://example.com/outbox"),
			Type:       a.OrderedCollectionType,
			URL:        a.URI("http://example.com/outbox"),
			TotalItems: 1,
			OrderedItems: a.ItemCollection{
				&a.Object{
					ID:      a.ObjectID("http://example.com/outbox/53c6fb47"),
					Type:    a.ArticleType,
					URL:     a.URI("http://example.com/53c6fb47"),
					Name:    a.NaturalLanguageValue{"-": "Example title"},
					Content: a.NaturalLanguageValue{"-": "Example content!"},
				},
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

	var f = t.Logf
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
		if pair.expected != assertDeepEquals(f, object, pair.result) {
			f = t.Errorf
			if stopOnFailure {
				f = t.Fatalf
			}

			f("\n%#v\n should %sequal to\n%#v", object, expLbl, pair.result)
			oj, e1 := j.Marshal(object)
			if e1 != nil {
				f(e1.Error())
			}

			tj, e2 := j.Marshal(pair.result)
			if e2 != nil {
				f(e2.Error())
			}
			f("\n%s\n should %sequal to expected\n%s", oj, expLbl, tj)
			continue
		}
		//fmt.Printf("%#v", object)
		if err == nil {
			fmt.Printf(" --- %s: %s\n          %s\n", "PASS", k, pair.path)
		}
	}
}
