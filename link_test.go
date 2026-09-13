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
		Rel       IRI
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
		Rel       IRI
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
	// link1 is a struct initialized inline which can be operated on directly.
	// For example, we can set the language:
	link1 := Link{Href: "http://example.com/1"}
	link1.HrefLang = English
	fmt.Printf("Link1: %v\n", link1)

	// link2 is wrapped in an Item interface.
	// It is probably the most common way of interacting with objects in the library,
	// because usually they get unmarshaled from an HTTP request, or another representation.
	var link2 Item = &Link{Type: LinkType}

	// That means we can't set any properties directly, so
	// if you uncommment the next line you will get a compiler error.
	// link2.Href = "http://example.com"

	_ = OnLink(link2, func(link *Link) error {
		// In order to operate on it, we must wrap it in a call to OnLink
		link.Href = "http://example.com/2"
		return nil
	})
	fmt.Printf("Link2: %v\n", link2)

	// Output:
	// Link1: activitypub.Link { href: http://example.com/1, hrefLang: en }
	// Link2: activitypub.Link[Link] { href: http://example.com/2 }
}
