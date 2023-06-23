package tests

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
	"unsafe"

	pub "github.com/go-ap/activitypub"

	j "github.com/go-ap/jsonld"
)

const dir = "./mocks"

var stopOnFailure = false

type testPair struct {
	expected bool
	blank    interface{}
	result   interface{}
}

type testMaps map[string]testPair

type visit struct {
	a1  unsafe.Pointer
	a2  unsafe.Pointer
	typ reflect.Type
}

type canErrorFunc func(format string, args ...interface{})

// See reflect.DeepEqual
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

// See reflect.deepValueEqual
func deepValueEqual(t canErrorFunc, v1, v2 reflect.Value, visited map[visit]bool, depth int) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.Type() != v2.Type() {
		t("types differ %s != %s", v1.Type().Name(), v2.Type().Name())
		return false
	}

	hard := func(v1, v2 reflect.Value) bool {
		switch v1.Kind() {
		case reflect.Ptr:
			return false
		case reflect.Map, reflect.Slice, reflect.Interface:
			// Nil pointers cannot be cyclic. Avoid putting them in the visited map.
			return !v1.IsNil() && !v2.IsNil()
		}
		return false
	}

	if hard(v1, v2) {
		var addr1, addr2 unsafe.Pointer
		if v1.CanAddr() {
			addr1 = unsafe.Pointer(v1.UnsafeAddr())
		} else {
			addr1 = unsafe.Pointer(v1.Pointer())
		}
		if v2.CanAddr() {
			addr2 = unsafe.Pointer(v2.UnsafeAddr())
		} else {
			addr2 = unsafe.Pointer(v2.Pointer())
		}
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
				isNil1 = "is not"
			}
			if v2.IsNil() {
				isNil2 = "is"
			} else {
				isNil2 = "is not"
			}
			t("Interface '%s' %s nil and '%s' %s nil", v1.Type().Name(), isNil1, v2.Type().Name(), isNil2)
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
			var (
				f1 = v1.Field(i)
				f2 = v2.Field(i)
				n1 = v1.Type().Field(i).Name
				n2 = v2.Type().Field(i).Name
				t1 = f1.Type().Name()
				t2 = f2.Type().Name()
			)
			if !deepValueEqual(t, v1.Field(i), v2.Field(i), visited, depth+1) {
				t("Struct fields at pos %d %s[%s] and %s[%s] are not deeply equal", i, n1, t1, n2, t2)
				if f1.CanInterface() && f2.CanInterface() {
					t("  Values: %#v - %#v", v1.Field(i).Interface(), v2.Field(i).Interface())
				}
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
	case reflect.String:
		return v1.String() == v2.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v1.Int() == v2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v1.Uint() == v2.Uint()
	case reflect.Float32, reflect.Float64:
		return v1.Float() == v2.Float()
	case reflect.Bool:
		return v1.Bool() == v2.Bool()
	case reflect.Complex64, reflect.Complex128:
		return v1.Complex() == v2.Complex()
	}
	return false
}

var zLoc, _ = time.LoadLocation("UTC")

var allTests = testMaps{
	//"empty": testPair{
	//	expected: true,
	//	blank:    &pub.Object{},
	//	result:   &pub.Object{},
	//},
	//"link_simple": testPair{
	//	expected: true,
	//	blank:    &pub.Link{},
	//	result: &pub.Link{
	//		Type:      pub.LinkType,
	//		Href:      pub.IRI("http://example.org/abc"),
	//		HrefLang:  pub.LangRef("en"),
	//		MediaType: pub.MimeType("text/html"),
	//		Name: pub.NaturalLanguageValues{{
	//			pub.NilLangRef, pub.Content("An example link"),
	//		}},
	//	},
	//},
	//"object_with_url": testPair{
	//	expected: true,
	//	blank:    &pub.Object{},
	//	result: &pub.Object{
	//		URL: pub.IRI("http://littr.git/api/accounts/system"),
	//	},
	//},
	//"object_with_url_collection": testPair{
	//	expected: true,
	//	blank:    &pub.Object{},
	//	result: &pub.Object{
	//		URL: pub.ItemCollection{
	//			pub.IRI("http://littr.git/api/accounts/system"),
	//			pub.IRI("http://littr.git/~system"),
	//		},
	//	},
	//},
	//"object_simple": testPair{
	//	expected: true,
	//	blank:    &pub.Object{},
	//	result: &pub.Object{
	//		Type: pub.ObjectType,
	//		ID:   pub.ID("http://www.test.example/object/1"),
	//		Name: pub.NaturalLanguageValues{{
	//			pub.NilLangRef, pub.Content("A Simple, non-specific object"),
	//		}},
	//	},
	//},
	//"object_no_type": testPair{
	//	expected: true,
	//	blank:    &pub.Object{},
	//	result: &pub.Object{
	//		ID:   pub.ID("http://www.test.example/object/1"),
	//		Name: pub.NaturalLanguageValues{{
	//			pub.NilLangRef, pub.Content("A Simple, non-specific object without a type"),
	//		}},
	//	},
	//},
	//"object_with_tags": testPair{
	//	expected: true,
	//	blank:    &pub.Object{},
	//	result: &pub.Object{
	//		Type: pub.ObjectType,
	//		ID:   pub.ID("http://www.test.example/object/1"),
	//		Name: pub.NaturalLanguageValues{{
	//			pub.NilLangRef, pub.Content("A Simple, non-specific object"),
	//		}},
	//		Tag: pub.ItemCollection{
	//			&pub.Mention{
	//				Name: pub.NaturalLanguageValues{{
	//					pub.NilLangRef, pub.Content("#my_tag"),
	//				}},
	//				Type: pub.MentionType,
	//				ID:   pub.ID("http://example.com/tag/my_tag"),
	//			},
	//			&pub.Mention{
	//				Name: pub.NaturalLanguageValues{{
	//					pub.NilLangRef, pub.Content("@ana"),
	//				}},
	//				Type: pub.MentionType,
	//				ID:   pub.ID("http://example.com/users/ana"),
	//			},
	//		},
	//	},
	//},
	//"object_with_replies": testPair{
	//	expected: true,
	//	blank:    &pub.Object{},
	//	result: &pub.Object{
	//		Type: pub.ObjectType,
	//		ID:   pub.ID("http://www.test.example/object/1"),
	//		Replies: &pub.Collection{
	//			ID:         pub.ID("http://www.test.example/object/1/replies"),
	//			Type:       pub.CollectionType,
	//			TotalItems: 1,
	//			Items: pub.ItemCollection{
	//				&pub.Object{
	//					ID:   pub.ID("http://www.test.example/object/1/replies/2"),
	//					Type: pub.ArticleType,
	//					Name: pub.NaturalLanguageValues{{
	//						pub.NilLangRef, pub.Content("Example title"),
	//					}},
	//				},
	//			},
	//		},
	//	},
	//},
	//"activity_simple": testPair{
	//	expected: true,
	//	blank: &pub.Activity{
	//		Actor: &pub.Person{},
	//	},
	//	result: &pub.Activity{
	//		Type:    pub.ActivityType,
	//		Summary: pub.NaturalLanguageValues{{pub.NilLangRef, pub.Content("Sally did something to a note")}},
	//		Actor: &pub.Person{
	//			Type: pub.PersonType,
	//			Name: pub.NaturalLanguageValues{{pub.NilLangRef, pub.Content("Sally")}},
	//		},
	//		Object: &pub.Object{
	//			Type: pub.NoteType,
	//			Name: pub.NaturalLanguageValues{{pub.NilLangRef, pub.Content("A Note")}},
	//		},
	//	},
	//},
	//"person_with_outbox": testPair{
	//	expected: true,
	//	blank:    &pub.Person{},
	//	result: &pub.Person{
	//		ID:                pub.ID("http://example.com/accounts/ana"),
	//		Type:              pub.PersonType,
	//		Name:              pub.NaturalLanguageValues{{pub.NilLangRef, pub.Content("ana")}},
	//		PreferredUsername: pub.NaturalLanguageValues{{pub.NilLangRef, pub.Content("Ana")}},
	//		URL:               pub.IRI("http://example.com/accounts/ana"),
	//		Outbox: &pub.OrderedCollection{
	//			ID:   "http://example.com/accounts/ana/outbox",
	//			Type: pub.OrderedCollectionType,
	//			URL:  pub.IRI("http://example.com/outbox"),
	//		},
	//	},
	//},
	//"ordered_collection": testPair{
	//	expected: true,
	//	blank:    &pub.OrderedCollection{},
	//	result: &pub.OrderedCollection{
	//		ID:         pub.ID("http://example.com/outbox"),
	//		Type:       pub.OrderedCollectionType,
	//		URL:        pub.IRI("http://example.com/outbox"),
	//		TotalItems: 1,
	//		OrderedItems: pub.ItemCollection{
	//			&pub.Object{
	//				ID:           pub.ID("http://example.com/outbox/53c6fb47"),
	//				Type:         pub.ArticleType,
	//				Name:         pub.NaturalLanguageValues{{pub.NilLangRef, pub.Content("Example title")}},
	//				Content:      pub.NaturalLanguageValues{{pub.NilLangRef, pub.Content("Example content!")}},
	//				URL:          pub.IRI("http://example.com/53c6fb47"),
	//				MediaType:    pub.MimeType("text/markdown"),
	//				Published:    time.Date(2018, time.July, 5, 16, 46, 44, 0, zLoc),
	//				Generator:    pub.IRI("http://example.com"),
	//				AttributedTo: pub.IRI("http://example.com/accounts/alice"),
	//			},
	//		},
	//	},
	//},
	"ordered_collection_page": testPair{
		expected: true,
		blank:    &pub.OrderedCollectionPage{},
		result: &pub.OrderedCollectionPage{
			PartOf:     pub.IRI("http://example.com/outbox"),
			Next:       pub.IRI("http://example.com/outbox?page=3"),
			Prev:       pub.IRI("http://example.com/outbox?page=1"),
			ID:         pub.ID("http://example.com/outbox?page=2"),
			Type:       pub.OrderedCollectionPageType,
			URL:        pub.IRI("http://example.com/outbox?page=2"),
			Current:    pub.IRI("http://example.com/outbox?page=2"),
			TotalItems: 1,
			StartIndex: 100,
			OrderedItems: pub.ItemCollection{
				&pub.Object{
					ID:           pub.ID("http://example.com/outbox/53c6fb47"),
					Type:         pub.ArticleType,
					Name:         pub.NaturalLanguageValues{{Ref: pub.NilLangRef, Value: pub.Content("Example title")}},
					Content:      pub.NaturalLanguageValues{{Ref: pub.NilLangRef, Value: pub.Content("Example content!")}},
					URL:          pub.IRI("http://example.com/53c6fb47"),
					MediaType:    pub.MimeType("text/markdown"),
					Published:    time.Date(2018, time.July, 5, 16, 46, 44, 0, zLoc),
					Generator:    pub.IRI("http://example.com"),
					AttributedTo: pub.IRI("http://example.com/accounts/alice"),
				},
			},
		},
	},
	//"natural_language_values": {
	//	expected: true,
	//	blank:  &pub.NaturalLanguageValues{},
	//	result: &pub.NaturalLanguageValues{
	//		{
	//			pub.NilLangRef, pub.Content([]byte{'\n','\t', '\t', '\n'}),
	//		},
	//		{pub.LangRef("en"), pub.Content("Ana got apples ⓐ")},
	//		{pub.LangRef("fr"), pub.Content("Aná a des pommes ⒜")},
	//		{pub.LangRef("ro"), pub.Content("Ana are mere")},
	//	},
	//},
	//"activity_create_simple": {
	//	expected: true,
	//	blank:    &pub.Create{},
	//	result: &pub.Create{
	//		Type:  pub.CreateType,
	//		Actor: pub.IRI("https://littr.git/api/accounts/anonymous"),
	//		Object: &pub.Object{
	//			Type:         pub.NoteType,
	//			AttributedTo: pub.IRI("https://littr.git/api/accounts/anonymous"),
	//			InReplyTo:    pub.IRI("https://littr.git/api/accounts/system/outbox/7ca154ff"),
	//			Content:      pub.NaturalLanguageValues{{pub.NilLangRef, pub.Content("<p>Hello world</p>")}},
	//			To:           pub.ItemCollection{pub.IRI("https://www.w3.org/ns/activitystreams#Public")},
	//		},
	//	},
	//},
	//"activity_create_multiple_objects": {
	//	expected: true,
	//	blank:    &pub.Create{},
	//	result: &pub.Create{
	//		Type:  pub.CreateType,
	//		Actor: pub.IRI("https://littr.git/api/accounts/anonymous"),
	//		Object: pub.ItemCollection{
	//			&pub.Object{
	//				Type:         pub.NoteType,
	//				AttributedTo: pub.IRI("https://littr.git/api/accounts/anonymous"),
	//				InReplyTo:    pub.IRI("https://littr.git/api/accounts/system/outbox/7ca154ff"),
	//				Content:      pub.NaturalLanguageValues{{pub.NilLangRef, pub.Content("<p>Hello world</p>")}},
	//				To:           pub.ItemCollection{pub.IRI("https://www.w3.org/ns/activitystreams#Public")},
	//			},
	//			&pub.Article{
	//				Type: pub.ArticleType,
	//				ID:   pub.ID("http://www.test.example/article/1"),
	//				Name: pub.NaturalLanguageValues{
	//					{
	//						pub.NilLangRef,
	//						pub.Content("This someday will grow up to be an article"),
	//					},
	//				},
	//				InReplyTo: pub.ItemCollection{
	//					pub.IRI("http://www.test.example/object/1"),
	//					pub.IRI("http://www.test.example/object/778"),
	//				},
	//			},
	//		},
	//	},
	//},
	//"object_with_audience": testPair{
	//	expected: true,
	//	blank:    &pub.Object{},
	//	result: &pub.Object{
	//		Type: pub.ObjectType,
	//		ID:   pub.ID("http://www.test.example/object/1"),
	//		To: pub.ItemCollection{
	//			pub.IRI("https://www.w3.org/ns/activitystreams#Public"),
	//		},
	//		Bto: pub.ItemCollection{
	//			pub.IRI("http://example.com/sharedInbox"),
	//		},
	//		CC: pub.ItemCollection{
	//			pub.IRI("https://example.com/actors/ana"),
	//			pub.IRI("https://example.com/actors/bob"),
	//		},
	//		BCC: pub.ItemCollection{
	//			pub.IRI("https://darkside.cookie/actors/darthvader"),
	//		},
	//	},
	//},
	//"article_with_multiple_inreplyto": {
	//	expected: true,
	//	blank:    &pub.Article{},
	//	result: &pub.Article{
	//		Type: pub.ArticleType,
	//		ID:   pub.ID("http://www.test.example/article/1"),
	//		Name: pub.NaturalLanguageValues{
	//			{
	//				pub.NilLangRef,
	//				pub.Content("This someday will grow up to be an article"),
	//			},
	//		},
	//		InReplyTo: pub.ItemCollection{
	//			pub.IRI("http://www.test.example/object/1"),
	//			pub.IRI("http://www.test.example/object/778"),
	//		},
	//	},
	//},
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

func TestUnmarshal(t *testing.T) {
	var err error

	f := t.Errorf
	if len(allTests) == 0 {
		t.Skip("No tests found")
	}

	for k, pair := range allTests {
		path := filepath.Join(dir, fmt.Sprintf("%s.json", k))
		t.Run(path, func(t *testing.T) {
			var data []byte
			data, err = getFileContents(path)
			if err != nil {
				f("Error: %s for %s", err, path)
				return
			}
			object := pair.blank

			err = j.Unmarshal(data, object)
			if err != nil {
				f("Error: %s for %s", err, data)
				return
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

				f("Mock: %s: %s\n%#v\n should %sequal to expected\n%#v", k, path, object, expLbl, pair.result)
				return
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
				f("Mock: %s: %s\n%s\n should %sequal to expected\n%s", k, path, oj, expLbl, tj)
			}
			//if err == nil {
			//	fmt.Printf(" --- %s: %s\n          %s\n", "PASS", k, path)
			//}
		})
	}
}
