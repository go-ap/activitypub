package activitypub

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLink_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_GobEncode(t *testing.T) {
	type fields struct {
		ID        ID
		Type      Typer
		Name      NaturalLanguageValues
		Rel       string
		MediaType MimeType
		Height    uint
		Width     uint
		Preview   Item
		Href      IRI
		HrefLang  LangRef
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			fields:  fields{},
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Link{
				ID:        tt.fields.ID,
				Type:      tt.fields.Type,
				Name:      tt.fields.Name,
				Rel:       tt.fields.Rel,
				MediaType: tt.fields.MediaType,
				Height:    tt.fields.Height,
				Width:     tt.fields.Width,
				Preview:   tt.fields.Preview,
				Href:      tt.fields.Href,
				HrefLang:  tt.fields.HrefLang,
			}
			got, err := l.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GobEncode() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestLink_GobDecode(t *testing.T) {
	type fields struct {
		ID        ID
		Type      Typer
		Name      NaturalLanguageValues
		Rel       string
		MediaType MimeType
		Height    uint
		Width     uint
		Preview   Item
		Href      IRI
		HrefLang  LangRef
	}
	tests := []struct {
		name    string
		fields  fields
		data    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			fields:  fields{},
			data:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Link{
				ID:        tt.fields.ID,
				Type:      tt.fields.Type,
				Name:      tt.fields.Name,
				Rel:       tt.fields.Rel,
				MediaType: tt.fields.MediaType,
				Height:    tt.fields.Height,
				Width:     tt.fields.Width,
				Preview:   tt.fields.Preview,
				Href:      tt.fields.Href,
				HrefLang:  tt.fields.HrefLang,
			}
			if err := l.GobDecode(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func ExampleLink_initialization() {
	// link1 is a struct literal which can be operated on directly.
	// For example, we can set the language:
	link1 := Link{Href: "http://example.com/1"}
	link1.HrefLang = English
	fmt.Printf("Link1: %v\n", link1)

	// link2 is wrapped in an Item interface.
	// It is probably the most common way of interacting with objects in the library,
	// because usually they get unmarshaled from an HTTP request, or another representation.
	var link2 Item = &Link{Type: LinkType}

	// That means we can't set any properties directly, so
	// uncommenting the next line will trigger a compiler error.
	//link1.Href = "http://example.com"

	_ = OnLink(link2, func(link *Link) error {
		// In order to operate on it, we must wrap it in a call to OnLink()
		link.Href = "http://example.com/2"
		return nil
	})
	fmt.Printf("Link2: %v\n", link2)

	// Output:
	// Link1: activitypub.Link { href: http://example.com/1, hrefLang: en }
	// Link2: activitypub.Link[Link] { href: http://example.com/2 }
}

func ExampleToLink() {
	// Alongside the Object compatible data types, ActivityPub and this library accept
	// Link types, which can be operated on in similar ways as types that satisfy the
	// ObjectOrLink interface.
	//
	// This fact represents the cornerstone for the functionality of the library,
	// as it allows us to access their properties of objects even when we're not certain
	// of their actual data type.

	// For Link, and it's associated type alias Mention, we can use the ToLink() to get
	// at the wrapped value.
	var link1 Item = &Mention{Type: MentionType, Name: DefaultNaturalLanguage("Jane Doe")}
	l1, _ := ToLink(link1)
	l1.Href = "http://example.com/~jdoe"
	fmt.Printf("Link1: %v\n", link1)
	fmt.Printf("     : %v\n\n", l1)

	// If the type wrapped in the interface is only a struct literal, modifying
	// the return value will not affect the original, which might have unexpected side effects.
	link2 := Link{Href: "http://example.com/2", Type: LinkType, Name: DefaultNaturalLanguage("link2")}
	l2, _ := ToLink(link2)
	l2.Name = nil
	fmt.Printf("Link2: %v\n", link2)
	fmt.Printf("     : %v\n\n", l2)

	// Because the Link and Object types are disjoint, an Object type can't be converted to a Link type,
	// and trying results in an error.
	var notLink Item = &Object{Type: LinkType, Name: DefaultNaturalLanguage("Eh?")}
	na, err := ToLink(notLink)
	fmt.Printf("NotLink: %v\n", notLink)
	fmt.Printf("       : %v\n", na)
	fmt.Printf("Error  : %v\n\n", err)

	// The reverse is also true, and you can see that in the ExampleToObject() test function.

	// Output:
	// Link1: activitypub.Link[Mention] { href: http://example.com/~jdoe, name: Jane Doe }
	//      : activitypub.Link[Mention] { href: http://example.com/~jdoe, name: Jane Doe }
	//
	// Link2: activitypub.Link[Link] { href: http://example.com/2, name: link2 }
	//      : activitypub.Link[Link] { href: http://example.com/2 }
	//
	// NotLink: activitypub.Object[Link] { name: Eh? }
	//        : <nil>
	// Error  : unable to convert *activitypub.Object to *activitypub.Link
}

func ExampleOnLink() {
	// link1 is wrapped in an Item interface.
	var link1 Item = &Link{Type: LinkType}

	// This means we can't set any properties directly, so
	// uncommenting the next line will trigger a compiler error:
	//link1.Href = "http://example.com"

	_ = OnLink(link1, func(link *Link) error {
		// In order to operate on it, we must wrap it in a call to OnLink()
		link.Href = "http://example.com"
		link.HrefLang = English

		// We can of course use the link's properties to read
		fmt.Printf("Link.Type: %v\n", link.Type)
		return nil
	})
	fmt.Printf("Link1: %v\n", link1)

	// Output:
	// Link.Type: Link
	// Link1: activitypub.Link[Link] { href: http://example.com, hrefLang: en }
}
