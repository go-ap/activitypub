package activitypub

import (
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

// A Link is an indirect, qualified reference to a resource identified by a URL.
// The fundamental model for links is established by [ RFC5988].
// Many of the properties defined by the Activity Vocabulary allow values that are either instances of APObject or Link.
// When a Link is used, it establishes a qualified relation connecting the subject
// (the containing object) to the resource identified by the href.
// Properties of the Link are properties of the reference as opposed to properties of the resource.
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
	write(&b, '{')

	if writeLinkJSONValue(&b, l) {
		write(&b, '}')
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
	return loadLink(val, l)
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (l *Link) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("UnmarshalBinary is not implemented for %T", *l)
}

/*
// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (l Link) MarshalBinary() ([]byte, error) {
	return nil, errors.New(fmt.Sprintf("MarshalBinary is not implemented for %T", l))
}

func (l Link) GobEncode() ([]byte, error) {
	return nil, errors.New(fmt.Sprintf("GobEncode is not implemented for %T", l))
}

func (l *Link) GobDecode([]byte) error {
	return errors.New(fmt.Sprintf("GobDecode is not implemented for %T", *l))
}
*/
