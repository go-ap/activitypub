package tests

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	pub "github.com/go-ap/activitypub"
	"github.com/google/go-cmp/cmp"

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
					Name:         pub.DefaultLangValue("Example title"),
					Content:      pub.DefaultLangValue("Example content!"),
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
	_, _ = io.ReadFull(f, data)
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
			status := cmp.Equal(object, pair.result)
			if pair.expected != status {
				if stopOnFailure {
					f = t.Fatalf
				}

				f("Mock: %s: %s\n should %sequal \n%s", k, path, expLbl, cmp.Diff(pair.result, object))
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
		})
	}
}
