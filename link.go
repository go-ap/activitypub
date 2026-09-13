package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/valyala/fastjson"
)

// LinkTypes represent the valid values for a Link object
var LinkTypes = ActivityVocabularyTypes{
	LinkType,
	MentionType,
}

type Links interface {
	Link | IRI
}

// Link represents an indirect, qualified reference to a resource identified by a URL.
// The fundamental model for links is established by [RFC5988].
// Many of the properties defined by the Activity-Vocabulary allow values that are either instances of Object or Link.
// When a Link is used, it establishes a qualified relation connecting the subject (the containing object) to the
// resource identified by the href. Properties of the Link are properties of the reference as opposed to properties
// of the resource.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-link
//
// Link describes a qualified, indirect reference to another resource that is closely related to the conceptual model
// of Links as established in [RFC5988]. The properties of the Link object are not the properties of the referenced
// resource, but are provided as hints for rendering agents to understand how to make use of the resource.
// For example, height and width might represent the desired rendered size of a referenced image, rather than the
// actual pixel dimensions of the referenced image.
// The target URI of the Link is expressed using the required href property. In addition, all Link instances share the
// following common set of optional properties as normatively defined by the [Activity Vocabulary]:
//
//	id | name | hreflang | mediaType | rel | height | width
//
// For example, all Objects can contain an image property whose value describes a graphical representation of the
// containing object. This property will typically be used to provide the URL to an image (e.g. JPEG, GIF or PNG)
// resource that can be displayed to the user. Any given object might have multiple such visual representations -
// multiple screenshots, for instance, or the same image at different resolutions. In [Activity Streams 2.0],
// there are essentially three ways of describing such references: a direct IRI, an embedded Object, an array of values
// which can be IRIs or Objects.
//
// https://www.w3.org/TR/activitystreams-core/#link
type Link struct {
	// Provides the globally unique identifier for an APObject or Link.
	ID ID `jsonld:"id,omitempty"`
	// Identifies the APObject or Link type. Multiple values may be specified.
	Type Typer `jsonld:"type,omitempty"`
	// A simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValues `jsonld:"name,omitempty,collapsible"`
	// A link relation associated with a Link. The value must conform to both the [HTML5] and
	// [RFC5988](https://tools.ietf.org/html/rfc5988) "link relation" definitions.
	// In the [HTML5], any string not containing the "space" U+0020, "tab" (U+0009), "LF" (U+000A),
	// "FF" (U+000C), "CR" (U+000D) or "," (U+002C) characters can be used as a valid link relation.
	Rel IRI `jsonld:"rel,omitempty"`
	// When used on a Link, identifies the MIME media type of the referenced resource.
	MediaType MimeType `jsonld:"mediaType,omitempty"`
	// On a Link, specifies a hint as to the rendering height in device-independent pixels of the linked resource.
	Height uint `jsonld:"height,omitempty"`
	// On a Link, specifies a hint as to the rendering width in device-independent pixels of the linked resource.
	Width uint `jsonld:"width,omitempty"`
	// Identifies an entity that provides a preview of this object.
	Preview Item `jsonld:"preview,omitempty"`
	// The target resource pointed to by a Link.
	Href IRI `jsonld:"href,omitempty"`
	// Hints as to the language used by the target resource.
	// Value must be a [BCP47](https://tools.ietf.org/html/bcp47) Language-Tag.
	HrefLang LangRef `jsonld:"hrefLang,omitempty"`
}

// Mention is a specialized Link that represents a @mention.
type Mention = Link

// GetID returns the ID corresponding to the Link object
func (l Link) GetID() ID {
	return l.ID
}

// GetLink returns the IRI corresponding to the current Link
func (l Link) GetLink() IRI {
	return IRI(l.ID)
}

// GetType returns the Type corresponding to the Mention object
func (l Link) GetType() Typer {
	return l.Type
}

// Match returns whether the receiver matches the ActivityVocabularyType arguments.
func (l Link) Match(tt ...ActivityVocabularyType) bool {
	return ActivityVocabularyTypes(tt).Match(l.Type)
}

// MarshalJSON encodes the receiver object to a JSON document.
func (l Link) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	JSONWrite(&b, '{')

	if JSONWriteLinkValue(&b, l) {
		JSONWrite(&b, '}')
		return b.Bytes(), nil
	}
	return nil, nil
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (l *Link) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}

	return jsonLoadToLink(val, l)
}

// Equals verifies if our receiver Link is equals with the "with" Item
func (l Link) Equals(other Item) bool {
	otherLink, err := ToLink(other)
	if err != nil {
		return false
	}
	return l.equal(*otherLink)
}

// equal verifies if our receiver Link is equals with the "with" Link
func (l Link) equal(with Link) bool {
	if !l.ID.Equal(with.ID) {
		return false
	}
	if !TypesEqual(l.Type, with.Type) {
		return false
	}
	if l.HrefLang != with.HrefLang {
		return false
	}
	if !l.Href.Equal(with.Href) {
		return false
	}
	if l.Rel != with.Rel {
		return false
	}
	if !l.Name.Equal(with.Name) {
		return false
	}
	if l.Height != with.Height {
		return false
	}
	if l.Width != with.Width {
		return false
	}
	return ItemsEqual(l.Preview, with.Preview)
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (l *Link) UnmarshalBinary(data []byte) error {
	return l.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (l Link) MarshalBinary() ([]byte, error) {
	return l.GobEncode()
}

func (l Link) GobEncode() ([]byte, error) {
	mm := make(map[string][]byte)
	hasData, err := mapLinkProperties(mm, l)
	if err != nil {
		return nil, err
	}
	if !hasData {
		return []byte{}, nil
	}
	bb := bytes.Buffer{}
	g := gob.NewEncoder(&bb)
	if err := g.Encode(mm); err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

func (l *Link) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	return unmapLinkProperties(mm, l)
}

func fmtLinkProps(w io.Writer) func(*Link) error {
	return func(l *Link) error {
		var n int
		comma := func() {
			if n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}

		if len(l.ID) > 0 {
			n, _ = fmt.Fprintf(w, "ID: %s", l.ID)
		}
		if len(l.Href) > 0 {
			comma()
			n, _ = fmt.Fprintf(w, "href: %s", l.Href)
		}
		if l.HrefLang != Und {
			comma()
			n, _ = fmt.Fprintf(w, "hrefLang: %s", l.HrefLang)
		}
		if len(l.Name) > 0 {
			comma()
			n, _ = fmt.Fprintf(w, "name: %s", l.Name)
		}
		if len(l.MediaType) > 0 {
			comma()
			n, _ = fmt.Fprintf(w, "mediaType: %s", l.MediaType)
		}
		if l.Height > 0 {
			comma()
			n, _ = fmt.Fprintf(w, "height: %d", l.Height)
		}
		if l.Width > 0 {
			comma()
			n, _ = fmt.Fprintf(w, "width: %d", l.Width)
		}
		if !IsNil(l.Preview) {
			comma()
			n, _ = fmt.Fprintf(w, "preview: %s", l.Preview)
		}
		if len(l.Rel) > 0 {
			comma()
			n, _ = fmt.Fprintf(w, "rel: %s", l.Rel)
		}
		return nil
	}
}

func (l Link) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		iri := l.ID
		if l.Href != "" {
			iri = l.Href
		}
		_, _ = s.Write([]byte(iri))
	case 'v':
		if l.Type != nil {
			_, _ = fmt.Fprintf(s, "%T[%v] { ", l, l.Type)
		} else {
			_, _ = fmt.Fprintf(s, "%T { ", l)
		}
		fmtLinkProps(s)(&l)
		s.Write([]byte(" }"))
	}
}

// ToLink returns a Link pointer to the data in the current Item
func ToLink(it LinkOrIRI) (*Link, error) {
	switch i := it.(type) {
	case *Link:
		return i, nil
	case Link:
		return &i, nil
	}
	return nil, fmt.Errorf("unable to convert %T to %T", it, new(Link))
}

// WithLinkFn represents a function type that can be used as a parameter for OnLink helper function
type WithLinkFn func(*Link) error

// OnLink calls function fn on the "it" LinkOrIRI if it can be asserted to type *Link
//
// This function should be safe to use for all types with a structure compatible
// with the Link type
func OnLink(it LinkOrIRI, fn WithLinkFn) error {
	if IsNil(it) {
		return nil
	}
	lnkFn := func(it LinkOrIRI) error {
		lnk, err := ToLink(it)
		if err != nil {
			return err
		}
		return fn(lnk)
	}

	if !IsItemCollection(it) {
		return lnkFn(it)
	}
	return OnItemCollection(it, func(col *ItemCollection) error {
		for _, ob := range *col {
			if err := lnkFn(ob); err != nil {
				return err
			}
		}
		return nil
	})
}

func notEmptyLink(l *Link) bool {
	return len(l.ID) > 0 ||
		LinkTypes.Match(l.GetType()) ||
		len(l.MediaType) > 0 ||
		l.Preview != nil ||
		l.Name != nil ||
		len(l.Href) > 0 ||
		len(l.Rel) > 0 ||
		l.HrefLang.Valid() ||
		l.Height > 0 ||
		l.Width > 0
}
