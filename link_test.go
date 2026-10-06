package activitypub

import (
	"fmt"
	"testing"

	"github.com/go-ap/errors"
	"github.com/google/go-cmp/cmp"
)

func TestLink_GetID(t *testing.T) {
	tests := []struct {
		name string
		ID   ID
		want ID
	}{
		{
			name: "empty",
			ID:   "",
			want: "",
		},
		{
			name: "not empty",
			ID:   "http://example.com",
			want: "http://example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Link{ID: tt.ID}
			if got := l.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLink_GetLink(t *testing.T) {
	tests := []struct {
		name string
		ID   ID
		want IRI
	}{
		{
			name: "empty",
			ID:   "",
			want: "",
		},
		{
			name: "not empty",
			ID:   "http://example.com",
			want: "http://example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Link{ID: tt.ID}
			if got := l.GetLink(); got != tt.want {
				t.Errorf("GetLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLink_GetType(t *testing.T) {
	tests := []struct {
		name string
		typ  Typer
		want Typer
	}{
		{
			name: "empty",
		},
		{
			name: "mention type",
			typ:  MentionType,
			want: MentionType,
		},
		{
			name: "link, mention",
			typ:  ActivityVocabularyTypes{MentionType, LinkType},
			want: ActivityVocabularyTypes{MentionType, LinkType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Link{Type: tt.typ}
			if got := l.GetType(); !TypesEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
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
	var link1 Item = &Mention{Type: MentionType, Name: DefaultLangValue("Jane Doe")}
	l1, _ := ToLink(link1)
	l1.Href = "http://example.com/~jdoe"
	fmt.Printf("Link1: %v\n", link1)
	fmt.Printf("     : %v\n\n", l1)

	// If the type wrapped in the interface is only a struct literal, modifying
	// the return value will not affect the original, which might have unexpected side effects.
	link2 := Link{Href: "http://example.com/2", Type: LinkType, Name: DefaultLangValue("link2")}
	l2, _ := ToLink(link2)
	l2.Name = nil
	fmt.Printf("Link2: %v\n", link2)
	fmt.Printf("     : %v\n\n", l2)

	// Because the Link and Object types are disjoint, an Object type can't be converted to a Link type,
	// and trying results in an error.
	var notLink Item = &Object{Type: LinkType, Name: DefaultLangValue("Eh?")}
	na, err := ToLink(notLink)
	fmt.Printf("NotLink: %v\n", notLink)
	fmt.Printf("       : %v\n", na)
	fmt.Printf("Error  : %v\n\n", err)

	// The reverse is also true, and you can see that in the ToObject() example function.

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

func TestLink_equal(t *testing.T) {
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
		name   string
		fields fields
		with   Link
		want   bool
	}{
		{
			name:   "empties equal",
			fields: fields{},
			with:   Link{},
			want:   true,
		},
		{
			name: "filled equal",
			fields: fields{
				ID:        "http://example.com/id",
				Type:      MentionType,
				Name:      DefaultLangValue("test"),
				Rel:       "alt",
				MediaType: "text/plain",
				Height:    10,
				Width:     20,
				Preview:   &Object{ID: "http://example.com/preview"},
				Href:      "http://example.com/href",
				HrefLang:  English,
			},
			with: Link{
				ID:        "http://example.com/id",
				Type:      MentionType,
				Name:      DefaultLangValue("test"),
				Rel:       "alt",
				MediaType: "text/plain",
				Height:    10,
				Width:     20,
				Preview:   &Object{ID: "http://example.com/preview"},
				Href:      "http://example.com/href",
				HrefLang:  English,
			},
			want: true,
		},
		{
			name:   "different types",
			fields: fields{Type: MentionType},
			with:   Link{Type: LinkType},
			want:   false,
		},
		{
			name:   "different id",
			fields: fields{ID: "http://example.com/different"},
			with:   Link{ID: "http://example.com/id"},
			want:   false,
		},
		{
			name:   "different name",
			fields: fields{Name: DefaultLangValue("test")},
			with:   Link{Name: DefaultLangValue("different")},
			want:   false,
		},
		{
			name:   "different rel",
			fields: fields{Rel: "alt"},
			with:   Link{Rel: "different"},
			want:   false,
		},
		{
			name:   "different mediaType",
			fields: fields{MediaType: "text/plain"},
			with:   Link{MediaType: "application/json"},
			want:   false,
		},
		{
			name:   "different height",
			fields: fields{Height: 666},
			with:   Link{Height: 10},
			want:   false,
		},
		{
			name:   "different width",
			fields: fields{Width: 20},
			with:   Link{Width: 21},
			want:   false,
		},
		{
			name:   "different preview",
			fields: fields{Preview: &Object{ID: "http://example.com/preview"}},
			with:   Link{Preview: IRI("http://example.com/preview")},
			want:   false,
		},
		{
			name:   "different href",
			fields: fields{Href: "http://example.com/href"},
			with:   Link{Href: "http://example.com/different"},
			want:   false,
		},
		{
			name:   "different hrefLang",
			fields: fields{HrefLang: English},
			with:   Link{HrefLang: French},
			want:   false,
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
			if got := l.equal(tt.with); got != tt.want {
				t.Errorf("equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLink_Match(t *testing.T) {
	tests := []struct {
		name string
		Type Typer
		with []ActivityVocabularyType
		want bool
	}{
		{
			name: "empty types match empty types",
			want: true,
		},
		{
			name: "link",
			Type: LinkType,
			with: ActivityVocabularyTypes{LinkType},
			want: true,
		},
		{
			name: "mention",
			Type: MentionType,
			with: ActivityVocabularyTypes{MentionType},
			want: true,
		},
		{
			name: "link not match mention",
			Type: LinkType,
			with: ActivityVocabularyTypes{MentionType},
			want: false,
		},
		{
			name: "link match multiple types",
			Type: LinkType,
			with: ActivityVocabularyTypes{LinkType, MentionType},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Link{Type: tt.Type}
			if got := l.Match(tt.with...); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOnLink(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		testFn  func(*testing.T) WithLinkFn
		wantErr error
	}{
		{
			name: "nil",
		},
		{
			name: "IRI is not a link",
			it:   IRI("http://example.com/m1"),
			testFn: func(t *testing.T) WithLinkFn {
				return func(l *Link) error {
					t.Errorf("unexpected to reach execution of link function")
					return nil
				}
			},
			wantErr: ErrorInvalidType[Link](IRI("")),
		},
		{
			name: "IRIs are not a link",
			it:   IRIs{"http://example.com/m1"},
			testFn: func(t *testing.T) WithLinkFn {
				return func(l *Link) error {
					t.Errorf("unexpected to reach execution of link function")
					return nil
				}
			},
			wantErr: ErrorInvalidType[Link](IRI("")),
		},
		{
			name: "IRI in ItemCollection is not a link",
			it:   ItemCollection{IRI("http://example.com/m1")},
			testFn: func(t *testing.T) WithLinkFn {
				return func(l *Link) error {
					t.Errorf("unexpected to reach execution of link function")
					return nil
				}
			},
			wantErr: ErrorInvalidType[Link](IRI("")),
		},
		{
			name: "object is not a link",
			it:   &Object{ID: "http://example.com/m1", Type: MentionType},
			testFn: func(t *testing.T) WithLinkFn {
				return func(l *Link) error {
					t.Errorf("unexpected to reach execution of link function")
					return nil
				}
			},
			wantErr: ErrorInvalidType[Link](&Object{}),
		},
		{
			name: "single link",
			it:   &Link{ID: "http://example.com/m1", Type: MentionType},
			testFn: func(t *testing.T) WithLinkFn {
				return func(l *Link) error {
					if !l.ID.Equal("http://example.com/m1") {
						t.Errorf("unexpected ID for link: %s, wanted %s", l.ID, "http://example.com/m1")
					}
					if l.Type != MentionType {
						t.Errorf("unexpected Type for link: %s, wanted %s", l.Type, MentionType)
					}
					return nil
				}
			},
		},
		{
			name: "links in item collection",
			it:   ItemCollection{Link{ID: "http://example.com/1"}, Link{ID: "http://example.com/2"}},
			testFn: func(t *testing.T) WithLinkFn {
				cnt := 0
				return func(link *Link) error {
					defer func() { cnt++ }()
					if cnt == 0 {
						if !link.ID.Equal("http://example.com/1") {
							t.Errorf("unexpected ID for first link in collection: %s, wanted %s", link.ID, "http://example.com/1")
						}
					}
					if cnt == 1 {
						if !link.ID.Equal("http://example.com/2") {
							t.Errorf("unexpected ID for first link in collection: %s, wanted %s", link.ID, "http://example.com/2")
						}
					}
					return nil
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(*Link) error { return nil }
			if tt.testFn != nil {
				fn = tt.testFn(t)
			}
			if err := OnLink(tt.it, fn); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("OnLink() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
		})
	}
}

func TestLink_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Link
		wantErr error
	}{
		{
			name:    "nil",
			data:    nil,
			wantErr: errors.NotImplementedf("Binary functionality not implemented"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Link{}
			if err := l.UnmarshalBinary(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalBinary() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestLink_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		sub     Link
		want    []byte
		wantErr error
	}{
		{
			name:    "nil",
			sub:     Link{},
			want:    nil,
			wantErr: errors.NotImplementedf("Binary functionality not implemented"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sub.MarshalBinary()
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("MarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("MarshalBinary() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
