package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"slices"

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
// resource that can be displayed to the user. Any given object might have multiple such visual representations --
// multiple screenshots, for instance, or the same image at different resolutions. In [Activity Streams 2.0],
// there are essentially three ways of describing such references.
//
// https://www.w3.org/TR/activitystreams-core/#link
type Link struct {
	// Provides the globally unique identifier for an APObject or Link.
	ID ID `jsonld:"id,omitempty"`
	// Identifies the APObject or Link type. Multiple values may be specified.
	Type ActivityVocabularyTypes `jsonld:"type,omitempty"`
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

// LinkNew initializes a new Link
func LinkNew(id ID, typ ActivityVocabularyType) *Link {
	if !LinkTypes.Contains(typ) {
		typ = LinkType
	}
	return &Link{ID: id, Type: typ.ToTypes()}
}

// MentionNew initializes a new Mention
func MentionNew(id ID) *Mention {
	return &Mention{ID: id, Type: MentionType.ToTypes()}
}

// IsLink validates if current Link is a Link
func (l Link) IsLink() bool {
	return l.GetType() == LinkType || LinkTypes.Contains(l.GetType())
}

// IsObject validates if current Link is an GetID
func (l Link) IsObject() bool {
	return l.GetType() == ObjectType || ObjectTypes.Contains(l.GetType())
}

// IsCollection returns false for Link objects
func (l Link) IsCollection() bool {
	return false
}

// GetID returns the ID corresponding to the Link object
func (l Link) GetID() ID {
	return l.ID
}

// GetLink returns the IRI corresponding to the current Link
func (l Link) GetLink() IRI {
	return IRI(l.ID)
}

// GetType returns the Type corresponding to the Mention object
func (l Link) GetType() ActivityVocabularyType {
	return l.Type.GetType()
}

// GetTypes returns the Types corresponding to the Mention object
func (l Link) GetTypes() ActivityVocabularyTypes {
	return l.Type
}

// MarshalJSON encodes the receiver object to a JSON document.
func (l Link) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	JSONWrite(&b, '{')

	if JSONWriteLinkValue(&b, l) {
		JSONWrite(&b, '}')
		return b, nil
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
func (l *Link) Equals(other Item) bool {
	if l == nil {
		return IsNil(other)
	}
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
	if slices.Equal(l.Type, with.Type) {
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
	if !ItemsEqual(l.Preview, with.Preview) {
		return false
	}
	return true
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

func (l Link) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		_, _ = fmt.Fprintf(s, "%T[%s] { %s }", l, l.Type, l.Href)
	}
}

// OnLink calls function fn on the "it" LinkOrIRI if it can be asserted to type *Link
//
// This function should be safe to use for all types with a structure compatible
// with the Link type
func OnLink(it LinkOrIRI, fn func(*Link) error) error {
	if it == nil {
		return nil
	}
	if IsItemCollection(it) {
		return callOnItemCollection(it, OnLink, fn)
	}
	ob, err := ToLink(it)
	if err != nil {
		return err
	}
	return fn(ob)
}
