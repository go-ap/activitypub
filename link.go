package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"

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
//   id | name | hreflang | mediaType | rel | height | width
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
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
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

// Mention is a specialized Link that represents an @mention.
type Mention = Link

// LinkNew initializes a new Link
func LinkNew(id ID, typ ActivityVocabularyType) *Link {
	if !LinkTypes.Contains(typ) {
		typ = LinkType
	}
	return &Link{ID: id, Type: typ}
}

// MentionNew initializes a new Mention
func MentionNew(id ID) *Mention {
	return &Mention{ID: id, Type: MentionType}
}

// IsLink validates if current Link is a Link
func (l Link) IsLink() bool {
	return l.Type == LinkType || LinkTypes.Contains(l.Type)
}

// IsObject validates if current Link is an GetID
func (l Link) IsObject() bool {
	return l.Type == ObjectType || ObjectTypes.Contains(l.Type)
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
		_, _ = fmt.Fprintf(s, "%T[%s] {  }", l, l.Type)
	}
}
