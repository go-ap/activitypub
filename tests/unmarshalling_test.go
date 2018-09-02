package tests

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
	"unsafe"

	a "github.com/mariusor/activitypub.go/activitypub"
	j "github.com/mariusor/activitypub.go/jsonld"
)

const dir = "./mocks"

var stopOnFailure = false

type testPair struct {
	path     string
	expected bool
	blank    interface{}
	result   interface{}
}

type tests map[string]testPair

type visit struct {
	a1  unsafe.Pointer
	a2  unsafe.Pointer
	typ reflect.Type
}

type canErrorFunc func(format string, args ...interface{})

func assertDeepEquals(t canErrorFunc, x, y interface{}) bool {
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

func deepValueEqual(t canErrorFunc, v1, v2 reflect.Value, visited map[visit]bool, depth int) bool {
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
			if v1.IsNil() == v2.IsNil() {
				return true
			}
			var isNil1, isNil2 string
			if v1.IsNil() {
				isNil1 = "is"
			} else {
				isNil1 = "isn't"
			}
			if v2.IsNil() {
				isNil2 = "is"
			} else {
				isNil2 = "isn't"
			}
			t("Interface %q %s nil and %q %s nil", v1.Type().Name(), isNil1, v2.Type().Name(), isNil2)
			return false
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
				t("Struct fields at pos %d %q:%q and %q:%q are not deeply equal", i, v1.Type().Field(i).Name, v1.Field(i).Type().Name(), v2.Type().Field(i).Name, v2.Field(i).Type().Name())
				t("  Values: %#v - %#v", v1.Field(i).Interface(), v2.Field(i).Interface())
				return false
			}
		}
		return true
	case reflect.Map:
		if v1.IsNil() != v2.IsNil() {
			t("Maps are not nil", v1.Type().Name(), v2.Type().Name())
			return false
		}
		if v1.Len() != v2.Len() {
			t("Maps don't have the same length %d vs %d", v1.Len(), v2.Len())
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !deepValueEqual(t, v1.MapIndex(k), v2.MapIndex(k), visited, depth+1) {
				t("Maps values at index %s are not equal", k.String())
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

var zLoc, _ = time.LoadLocation("UTC")

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
			Name: a.NaturalLanguageValue{{
				a.NilLangRef, "An example link",
			}},
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
			Name: a.NaturalLanguageValue{{
				a.NilLangRef, "A Simple, non-specific object",
			}},
		},
	},
	"object_with_replies": testPair{
		path:     "./mocks/object_with_replies.json",
		expected: true,
		blank:    &a.Object{},
		result: &a.Object{
			Type: a.ObjectType,
			ID:   a.ObjectID("http://www.test.example/object/1"),
			Replies: &a.Collection{
				ID:         a.ObjectID("http://www.test.example/object/1/replies"),
				Type:       a.CollectionType,
				TotalItems: 1,
				Items: a.ItemCollection{
					&a.Object{
						ID:   a.ObjectID("http://www.test.example/object/1/replies/2"),
						Type: a.ArticleType,
						Name: a.NaturalLanguageValue{{a.NilLangRef, "Example title"}},
					},
				},
			},
		},
	},
	//"activity_simple": testPair{
	//	path:     "./mocks/activity_simple.json",
	//	expected: false,
	//	blank:    &a.Activity{},
	//	result: &a.Activity{
	//		Type:    a.ActivityType,
	//		Summary: a.NaturalLanguageValue{{a.NilLangRef, "Sally did something to a note"}},
	//		Actor: a.Actor(a.Person{
	//			Type: a.PersonType,
	//			Name: a.NaturalLanguageValue{{
	//				a.NilLangRef, "Sally",
	//			}},
	//		}),
	//		Object: a.Object{
	//			Type: a.NoteType,
	//			Name: a.NaturalLanguageValue{{
	//				a.NilLangRef, "A Note",
	//			}},
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
			Name: a.NaturalLanguageValue{{
				a.NilLangRef, "ana",
			}},
			PreferredUsername: a.NaturalLanguageValue{{
				a.NilLangRef, "Ana",
			}},
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
		expected: true,
		blank:    &a.OrderedCollection{},
		result: &a.OrderedCollection{
			ID:         a.ObjectID("http://example.com/outbox"),
			Type:       a.OrderedCollectionType,
			URL:        a.URI("http://example.com/outbox"),
			TotalItems: 1,
			OrderedItems: a.ItemCollection{
				&a.Object{
					ID:           a.ObjectID("http://example.com/outbox/53c6fb47"),
					Type:         a.ArticleType,
					Name:         a.NaturalLanguageValue{{a.NilLangRef, "Example title"}},
					Content:      a.NaturalLanguageValue{{a.NilLangRef, "Example content!"}},
					URL:          a.URI("http://example.com/53c6fb47"),
					MediaType:    a.MimeType("text/markdown"),
					Published:    time.Date(2018, time.July, 5, 16, 46, 44, 0, zLoc),
					Generator:    a.IRI("http://example.com"),
					AttributedTo: a.IRI("http://example.com/accounts/alice"),
				},
			},
		},
	},
	"natural_language_values": {
		path:     "./mocks/natural_language_values.json",
		expected: true,
		blank:    &a.NaturalLanguageValue{},
		result: &a.NaturalLanguageValue{
			{
				a.NilLangRef, `
	
	`},
			{a.LangRef("en"), "Ana got apples ⓐ"},
			{a.LangRef("fr"), "Aná a des pommes ⒜"},
			{a.LangRef("ro"), "Ana are mere"},
		},
	},
	"create_activity_simple": {
		path:     "./mocks/create_activity_simple.json",
		expected: true,
		blank:    &a.Create{},
		result: &a.Create{
			Type:  a.CreateType,
			Actor: a.IRI("https://littr.git/api/accounts/anonymous"),
			Object: &a.Object{
				Type:         a.NoteType,
				AttributedTo: a.IRI("https://littr.git/api/accounts/anonymous"),
				InReplyTo:    a.IRI("https://littr.git/api/accounts/system/outbox/7ca154ff"),
				Content:      a.NaturalLanguageValue{{a.NilLangRef, "<p>Hello world</p>"}},
				To:           a.ObjectsArr{a.IRI("https://www.w3.org/ns/activitystreams#Public")},
			},
		},
	},
}

func getFileContents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	st, err := f.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, st.Size())
	io.ReadFull(f, data)
	data = bytes.Trim(data, "\x00")

	return data, nil
}

func Test_ActivityPubUnmarshall(t *testing.T) {
	var err error

	var f = t.Errorf
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
		status := assertDeepEquals(f, object, pair.result)
		if pair.expected != status {
			if stopOnFailure {
				f = t.Fatalf
			}

			f("\n%#v\n should %sequal to expected\n%#v", object, expLbl, pair.result)
			continue
		}
		if !status {
			oj, err := j.Marshal(object)
			if err != nil {
				f(err.Error())
			}
			tj, err := j.Marshal(pair.result)
			if err != nil {
				f(err.Error())
			}
			f("\n%s\n should %sequal to expected\n%s", oj, expLbl, tj)
		}
		if err == nil {
			fmt.Printf(" --- %s: %s\n          %s\n", "PASS", k, pair.path)
		}
	}
}
